package listeners

import (
	"context"
	"sync"
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	txschedulertypes "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/txscheduler"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	txscheduler "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/client"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-sentry/tx-sentry/use-cases"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -source=session_manager.go -destination=mocks/session_manager.go -package=mocks

const sessionManagerComponent = "service.session-manager"

type SessionManager interface {
	Start(ctx context.Context, job *entities.Job)
}

// sessionManager is a manager of job sessions
type sessionManager struct {
	mutex                  *sync.RWMutex
	sessions               map[string]bool
	retrySessionJobUseCase usecases.RetrySessionJobUseCase
	txSchedulerClient      txscheduler.TransactionSchedulerClient
}

type sessionData struct {
	parentJob        *entities.Job
	nChildren        int
	retries          int
	lastChildJobUUID string
}

// NewSessionManager creates a new SessionManager
func NewSessionManager(txSchedulerClient txscheduler.TransactionSchedulerClient, retrySessionJobUseCase usecases.RetrySessionJobUseCase) SessionManager {
	return &sessionManager{
		mutex:                  &sync.RWMutex{},
		sessions:               make(map[string]bool),
		retrySessionJobUseCase: retrySessionJobUseCase,
		txSchedulerClient:      txSchedulerClient,
	}
}

func (manager *sessionManager) Start(ctx context.Context, job *entities.Job) {
	logger := log.WithContext(ctx).WithField("job_uuid", job.UUID)

	if manager.hasSession(job.UUID) {
		logger.Debug("job session already exists, skipping session creation")
		return
	}

	if job.InternalData.RetryInterval == 0 {
		logger.Debug("job session has retry strategy disabled")
		return
	}

	if job.InternalData.HasBeenRetried {
		logger.Warn("job session been already retried")
		return
	}

	ses, err := manager.retrieveJobSessionData(ctx, job)
	if err != nil {
		logger.WithError(err).Error("job listening session failed to start")
		return
	}

	if ses.retries >= txschedulertypes.SentryMaxRetries {
		logger.Warn("job already reached max retries")
		return
	}

	manager.addSession(job.UUID)

	go func() {
		bckOff := backoff.NewExponentialBackOff()
		bckOff.MaxInterval = 5 * time.Second
		bckOff.MaxElapsedTime = time.Minute
		err := backoff.RetryNotify(
			func() error {
				err := manager.runSession(ctx, ses)
				return err
			},
			bckOff,
			func(err error, d time.Duration) {
				logger.WithError(err).Warnf("error in job retry session, restarting in %v...", d)
			},
		)

		if err != nil {
			logger.WithError(err).Error("job listening session unexpectedly stopped")
		}

		annotations := txschedulertypes.FormatInternalDataToAnnotations(job.InternalData)
		annotations.HasBeenRetried = true
		_, err = manager.txSchedulerClient.UpdateJob(ctx, job.UUID, &txschedulertypes.UpdateJobRequest{
			Annotations: &annotations,
		})

		if err != nil {
			logger.WithError(err).Error("failed to update job labels")
		}

		manager.removeSession(job.UUID)
	}()
}

func (manager *sessionManager) addSession(jobUUID string) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	manager.sessions[jobUUID] = true
}

func (manager *sessionManager) hasSession(jobUUID string) bool {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()
	_, ok := manager.sessions[jobUUID]
	return ok
}

func (manager *sessionManager) runSession(ctx context.Context, ses *sessionData) error {
	logger := log.WithContext(ctx).WithField("job_uuid", ses.parentJob.UUID)
	logger.Info("starting job session")

	ticker := time.NewTicker(ses.parentJob.InternalData.RetryInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			childJobUUID, err := manager.retrySessionJobUseCase.Execute(ctx, ses.parentJob.UUID, ses.lastChildJobUUID, ses.nChildren)
			if err != nil {
				return errors.FromError(err).ExtendComponent(sessionManagerComponent)
			}

			ses.retries++
			if ses.retries >= txschedulertypes.SentryMaxRetries {
				return nil
			}

			// If no child created but no error, we exit the session gracefully
			if childJobUUID == "" {
				return nil
			}

			if childJobUUID != ses.lastChildJobUUID {
				ses.nChildren++
				ses.lastChildJobUUID = childJobUUID
			}
		case <-ctx.Done():
			logger.WithField("reason", ctx.Err().Error()).Info("session gracefully stopped")
			return nil
		}
	}
}

func (manager *sessionManager) removeSession(jobUUID string) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	delete(manager.sessions, jobUUID)
}

func (manager *sessionManager) retrieveJobSessionData(ctx context.Context, job *entities.Job) (*sessionData, error) {
	jobs, err := manager.txSchedulerClient.SearchJob(ctx, &entities.JobFilters{
		ChainUUID:     job.ChainUUID,
		ParentJobUUID: job.UUID,
	})

	if err != nil {
		return nil, err
	}

	nChildren := len(jobs) - 1
	lastJobRetry := jobs[len(jobs)-1]

	// we count the number of resending of last job as retries
	nRetries := nChildren
	for _, lg := range lastJobRetry.Logs {
		if lg.Status == utils.StatusResending {
			nRetries++
		}
	}

	return &sessionData{
		parentJob:        job,
		nChildren:        nChildren,
		retries:          nRetries,
		lastChildJobUUID: jobs[nChildren].UUID,
	}, nil
}
