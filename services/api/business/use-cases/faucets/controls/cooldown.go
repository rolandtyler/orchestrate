package controls

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/consensys/orchestrate/pkg/types/entities"

	"github.com/consensys/orchestrate/pkg/errors"
	stripedmutex "github.com/nmvalera/striped-mutex"
)

const cooldownComponent = "faucet.control.cooldown"

// Controller that forces a minimum time interval between 2 credits
type CooldownControl struct {
	mux            *stripedmutex.StripedMutex
	lastAuthorized *sync.Map
}

// NewController creates a CoolDown controller
func NewCooldownControl() *CooldownControl {
	return &CooldownControl{
		lastAuthorized: &sync.Map{},
		mux:            stripedmutex.New(100),
	}
}

// Control apply CoolDown controller on a credit function
func (ctrl *CooldownControl) Control(ctx context.Context, req *entities.FaucetRequest) error {
	if len(req.Candidates) == 0 {
		return nil
	}

	// If still cooling down we invalid credit
	for key, candidate := range req.Candidates {
		ctrl.lock(key, req.Beneficiary.String())
		defer ctrl.unlock(key, req.Beneficiary.String())

		if ctrl.IsCoolingDown(key, req.Beneficiary.String(), candidate.Cooldown) {
			log.FromContext(ctx).Debug("candidate removed due to CooldownControl")
			delete(req.Candidates, key)
		}
	}

	if len(req.Candidates) == 0 {
		errMessage := "all faucets cooling down"
		log.FromContext(ctx).Error(errMessage)
		return errors.FaucetWarning(errMessage).ExtendComponent(cooldownComponent)
	}

	return nil
}

func (ctrl *CooldownControl) OnSelectedCandidate(_ context.Context, faucet *entities.Faucet, beneficiary ethcommon.Address) error {
	ctrl.lastAuthorized.Store(mapKey(faucet.UUID, beneficiary.String()), time.Now())
	return nil
}

// IsCoolingDown indicates if faucet is cooling down
func (ctrl *CooldownControl) IsCoolingDown(faucetID, beneficiary, cooldown string) bool {
	delay, _ := time.ParseDuration(cooldown)

	lastAuthorized, _ := ctrl.lastAuthorized.LoadOrStore(mapKey(faucetID, beneficiary), time.Time{})
	return time.Since(lastAuthorized.(time.Time)) < delay
}

func (ctrl *CooldownControl) lock(faucetID, account string) {
	ctrl.mux.Lock(mapKey(faucetID, account))
}

func (ctrl *CooldownControl) unlock(faucetID, account string) {
	ctrl.mux.Unlock(mapKey(faucetID, account))
}

func mapKey(faucetID, account string) string {
	return fmt.Sprintf("%s@%s", account, faucetID)
}
