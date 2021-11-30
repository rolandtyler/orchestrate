package nonce

import (
	"context"
	"crypto/md5"
	"fmt"
	"sort"
	"strings"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/toolkit/ethclient"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/types/tx"
	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/consensys/orchestrate/services/tx-sender/store"
)

const component = "nonce-manager"
const fetchNonceErr = "cannot retrieve fetch nonce from chain"

//go:generate mockgen -source=manager.go -destination=mocks/manager.go -package=mocks

type Manager interface {
	GetNonce(ctx context.Context, job *entities.Job) (uint64, error)
	CleanNonce(ctx context.Context, job *entities.Job, jobErr error) error
	IncrementNonce(ctx context.Context, job *entities.Job) error
}

type nonceManager struct {
	nonce            store.NonceSender
	ethClient        ethclient.MultiClient
	recovery         store.RecoveryTracker
	maxRecovery      uint64
	chainRegistryURL string
	logger           *log.Logger
}

func NewNonceManager(ec ethclient.MultiClient, nm store.NonceSender, tracker store.RecoveryTracker, chainRegistryURL string,
	maxRecovery uint64) Manager {
	return &nonceManager{
		nonce:            nm,
		ethClient:        ec,
		recovery:         tracker,
		maxRecovery:      maxRecovery,
		chainRegistryURL: chainRegistryURL,
		logger:           log.NewLogger().SetComponent(component),
	}
}

func (nc *nonceManager) GetNonce(ctx context.Context, job *entities.Job) (uint64, error) {
	logger := nc.logger.WithContext(ctx).WithField("job", job.UUID)

	nonceKey := partitionKey(job)
	if nonceKey == "" {
		logger.Debug("empty nonceKey, skip nonce check")
		return 0, nil
	}

	// Retrieve last sent nonce from nonce manager
	lastSent, ok, err := nc.nonce.GetLastSent(nonceKey)
	if err != nil {
		errMsg := "cannot retrieve lastSent nonce"
		logger.WithError(err).Error(errMsg)
		return 0, err
	}

	var expectedNonce uint64
	if ok {
		expectedNonce = lastSent + 1
	} else {
		expectedNonce, err = nc.fetchNonceFromChain(ctx, job)
		if err != nil {
			logger.WithError(err).Error(fetchNonceErr)
			return 0, err
		}

		logger.WithField("pending_nonce", expectedNonce).Debug("calibrating nonce")
	}

	return expectedNonce, nil
}

func (nc *nonceManager) CleanNonce(ctx context.Context, job *entities.Job, jobErr error) error {
	logger := nc.logger.WithContext(ctx).WithField("job", job.UUID)

	if job.InternalData.ParentJobUUID == job.UUID {
		logger.Debug("ignored nonce errors in children jobs")
		return nil
	}

	// TODO: update EthClient to process and standardize nonce too low errors
	if !strings.Contains(strings.ToLower(jobErr.Error()), "nonce too low") &&
		!strings.Contains(strings.ToLower(jobErr.Error()), "incorrect nonce") &&
		!strings.Contains(strings.ToLower(jobErr.Error()), "replacement transaction") {
		return nil
	}

	nonceKey := partitionKey(job)
	logger.Warn("chain responded with invalid nonce error")
	if nc.recovery.Recovering(job.UUID) >= nc.maxRecovery {
		err := errors.InvalidNonceError("reached max nonce recovery max")
		logger.WithError(err).Error("cannot recover from nonce error")
		return err
	}

	txNonce := uint64(0)
	if job.Transaction.Nonce != nil {
		txNonce = *job.Transaction.Nonce
	}

	// Clean nonce value only if it was used to set the txNonce
	lastSentNonce, ok, _ := nc.nonce.GetLastSent(nonceKey)
	if ok && txNonce == lastSentNonce+1 {
		logger.WithField("last_sent", lastSentNonce).Debug("cleaning account nonce")
		if err := nc.nonce.DeleteLastSent(nonceKey); err != nil {
			logger.WithError(err).Error("cannot clean NonceManager LastSent")
			return err
		}
	}

	// In case of failing because "nonce too low" we reset tx nonce
	nc.recovery.Recover(job.UUID)
	job.Transaction.Nonce = nil

	return errors.InvalidNonceWarning(jobErr.Error())
}

func (nc *nonceManager) IncrementNonce(ctx context.Context, job *entities.Job) error {
	logger := nc.logger.WithContext(ctx).WithField("job", job.UUID)

	nonceKey := partitionKey(job)
	txNonce := uint64(0)
	if job.Transaction.Nonce != nil {
		txNonce = *job.Transaction.Nonce
	}

	// Set nonce value only if txNonce was using previous value
	lastSentNonce, ok, _ := nc.nonce.GetLastSent(nonceKey)
	if !ok || txNonce == lastSentNonce+1 {
		err := nc.nonce.SetLastSent(nonceKey, txNonce)
		if err != nil {
			logger.WithError(err).Error("could not store last sent nonce")
			return err
		}
	}

	logger.WithField("last_sent", *job.Transaction.Nonce).Debug("increment account nonce value")
	nc.recovery.Recovered(job.UUID)
	return nil
}

func (nc *nonceManager) fetchNonceFromChain(ctx context.Context, job *entities.Job) (n uint64, err error) {
	url := utils.GetProxyURL(nc.chainRegistryURL, job.ChainUUID)

	switch {
	case string(job.Type) == tx.JobType_ETH_EEA_PRIVATE_TX.String() && job.Transaction.PrivacyGroupID != nil:
		n, err = nc.ethClient.PrivNonce(ctx, url, *job.Transaction.From,
			job.Transaction.PrivacyGroupID)
	case string(job.Type) == tx.JobType_ETH_EEA_PRIVATE_TX.String() && job.Transaction.PrivateFor != nil:
		n, err = nc.ethClient.PrivEEANonce(ctx, url, *job.Transaction.From,
			job.Transaction.PrivateFrom, job.Transaction.PrivateFor)
	default:
		n, err = nc.ethClient.PendingNonceAt(ctx, url, *job.Transaction.From)
	}

	return
}

func partitionKey(job *entities.Job) string {
	// Return empty partition key for raw tx and one time key tx
	// Not able to format a correct partition key if From or ChainID are not set. In that case return empty partition key
	if job.Transaction.From == nil || job.InternalData.ChainID == nil {
		return ""
	}

	fromAddr := job.Transaction.From
	chainID := job.InternalData.ChainID
	switch {
	case string(job.Type) == tx.JobType_ETH_EEA_PRIVATE_TX.String() && job.Transaction.PrivacyGroupID != nil:
		return fmt.Sprintf("%v@eea-%v@%v", fromAddr, job.Transaction.PrivacyGroupID, chainID)
	case string(job.Type) == tx.JobType_ETH_EEA_PRIVATE_TX.String() && len(job.Transaction.PrivateFor) > 0:
		l := utils.ArrBytesToString(append(job.Transaction.PrivateFor, job.Transaction.PrivateFrom))
		sort.Strings(l)
		h := md5.New()
		_, _ = h.Write([]byte(strings.Join(l, "-")))
		return fmt.Sprintf("%v@eea-%v@%v", fromAddr, fmt.Sprintf("%x", h.Sum(nil)), chainID)
	default:
		return fmt.Sprintf("%v@%v", fromAddr, chainID)
	}
}
