package faucets

import (
	"context"
	"reflect"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/toolkit/ethclient"
	"github.com/consensys/orchestrate/pkg/types/entities"
	usecases "github.com/consensys/orchestrate/services/api/business/use-cases"
	"github.com/consensys/orchestrate/services/api/business/use-cases/faucets/controls"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

const getFaucetCandidateComponent = "use-cases.faucet-candidate"

type FaucetControl interface {
	Control(ctx context.Context, req *entities.FaucetRequest) error
	OnSelectedCandidate(ctx context.Context, faucet *entities.Faucet, beneficiary ethcommon.Address) error
}

// RegisterContract is a use case to register a new contract
type faucetCandidate struct {
	chainStateReader ethclient.ChainStateReader
	searchFaucets    usecases.SearchFaucetsUseCase
	controls         []FaucetControl
	logger           *log.Logger
}

// NewGetFaucetCandidateUseCase creates a new GetFaucetCandidateUseCase
func NewGetFaucetCandidateUseCase(
	searchFaucets usecases.SearchFaucetsUseCase,
	chainStateReader ethclient.ChainStateReader,
) usecases.GetFaucetCandidateUseCase {
	cooldownCtrl := controls.NewCooldownControl()
	maxBalanceCtrl := controls.NewMaxBalanceControl(chainStateReader)
	creditorCtrl := controls.NewCreditorControl(chainStateReader)

	return &faucetCandidate{
		chainStateReader: chainStateReader,
		searchFaucets:    searchFaucets,
		controls:         []FaucetControl{creditorCtrl, cooldownCtrl, maxBalanceCtrl},
		logger:           log.NewLogger().SetComponent(getFaucetCandidateComponent),
	}
}

func (uc *faucetCandidate) Execute(ctx context.Context, account ethcommon.Address, chain *entities.Chain, userInfo *multitenancy.UserInfo) (*entities.Faucet, error) {
	ctx = log.With(log.WithFields(ctx, log.Field("chain", chain.UUID), log.Field("account", account)), uc.logger)
	logger := uc.logger.WithContext(ctx)

	faucets, err := uc.searchFaucets.Execute(ctx, &entities.FaucetFilters{ChainRule: chain.UUID}, userInfo)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(getFaucetCandidateComponent)
	}

	if len(faucets) == 0 {
		errMessage := "no faucet candidate found"
		logger.Debug(errMessage)
		return nil, errors.NotFoundError(errMessage).ExtendComponent(getFaucetCandidateComponent)
	}

	candidates := make(map[string]*entities.Faucet)
	for _, faucet := range faucets {
		candidates[faucet.UUID] = faucet
	}
	req := &entities.FaucetRequest{
		Beneficiary: account,
		Candidates:  candidates,
		Chain:       chain,
	}
	for _, ctrl := range uc.controls {
		err = ctrl.Control(ctx, req)
		if err != nil {
			return nil, errors.FromError(err).ExtendComponent(getFaucetCandidateComponent)
		}
	}

	if len(req.Candidates) == 0 {
		errMessage := "no faucet candidate retained"
		logger.Debug(errMessage)
		return nil, errors.NotFoundError(errMessage).ExtendComponent(getFaucetCandidateComponent)
	}

	// Select a first faucet candidate for comparison
	selectedFaucet := req.Candidates[electFaucet(req.Candidates)]
	for _, ctrl := range uc.controls {
		err := ctrl.OnSelectedCandidate(ctx, selectedFaucet, req.Beneficiary)
		if err != nil {
			return nil, errors.FromError(err).ExtendComponent(getFaucetCandidateComponent)
		}
	}

	logger.WithField("creditor_account", selectedFaucet.CreditorAccount).Debug("faucet candidate found successfully")
	return selectedFaucet, nil
}

// electFaucet is currently selecting the remaining faucet candidates with the highest amount
func electFaucet(faucetsCandidates map[string]*entities.Faucet) string {
	electedFaucet := reflect.ValueOf(faucetsCandidates).MapKeys()[0].String()
	amountElectedFaucet := faucetsCandidates[electedFaucet].Amount.ToInt()

	for key, candidate := range faucetsCandidates {
		if candidate.Amount.ToInt().Cmp(amountElectedFaucet) > 0 {
			electedFaucet = key
		}
	}

	return electedFaucet
}
