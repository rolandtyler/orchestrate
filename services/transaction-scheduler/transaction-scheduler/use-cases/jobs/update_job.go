package jobs

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/parsers"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases"
)

const updateJobComponent = "use-cases.update-job"

// updateJobUseCase is a use case to create a new transaction job
type updateJobUseCase struct {
	db                    store.DB
	updateChildrenUseCase usecases.UpdateChildrenUseCase
	startNextJobUseCase   usecases.StartNextJobUseCase
}

// NewUpdateJobUseCase creates a new UpdateJobUseCase
func NewUpdateJobUseCase(db store.DB, updateChildrenUseCase usecases.UpdateChildrenUseCase, startJobUC usecases.StartNextJobUseCase) usecases.UpdateJobUseCase {
	return &updateJobUseCase{
		db:                    db,
		updateChildrenUseCase: updateChildrenUseCase,
		startNextJobUseCase:   startJobUC,
	}
}

// Execute validates and creates a new transaction job
func (uc *updateJobUseCase) Execute(ctx context.Context, job *entities.Job, nextStatus, logMessage string, tenants []string) (*entities.Job, error) {
	logger := log.WithContext(ctx).WithField("tenants", tenants).WithField("job_uuid", job.UUID).
		WithField("next_status", nextStatus)
	logger.Debug("updating job entity")

	jobModel, err := uc.db.Job().FindOneByUUID(ctx, job.UUID, tenants, true)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(updateJobComponent)
	}

	if utils.IsFinalStatus(jobModel.Status) {
		errMessage := "job status is final, cannot be updated"
		logger.WithField("status", jobModel.Status).Error(errMessage)
		return nil, errors.InvalidParameterError(errMessage).ExtendComponent(updateJobComponent)
	}

	// We are not forced to update the transaction
	if job.Transaction != nil {
		parsers.UpdateTransactionModelFromEntities(jobModel.Transaction, job.Transaction)
		if err = uc.db.Transaction().Update(ctx, jobModel.Transaction); err != nil {
			return nil, errors.FromError(err).ExtendComponent(updateJobComponent)
		}
	}

	if len(job.Labels) > 0 {
		jobModel.Labels = job.Labels
	}
	if job.InternalData != nil {
		jobModel.InternalData = job.InternalData
	}

	var jobLogModel *models.Log
	// We are not forced to update the status
	if nextStatus != "" && !canUpdateStatus(nextStatus, jobModel.Status) {
		errMessage := "invalid status update for the current job state"
		logger.WithField("status", jobModel.Status).WithField("next_status", nextStatus).Error(errMessage)
		return nil, errors.InvalidStateError(errMessage).ExtendComponent(updateJobComponent)
	} else if nextStatus != "" {
		jobLogModel = &models.Log{
			JobID:   &jobModel.ID,
			Status:  nextStatus,
			Message: logMessage,
		}
	}

	// In case of status update
	err = uc.updateJob(ctx, jobModel, jobLogModel)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(updateJobComponent)
	}

	if (nextStatus == utils.StatusMined || nextStatus == utils.StatusStored) && jobModel.NextJobUUID != "" {
		err = uc.startNextJobUseCase.Execute(ctx, jobModel.UUID, tenants)
		if err != nil {
			return nil, errors.FromError(err).ExtendComponent(updateJobComponent)
		}
	}

	log.WithContext(ctx).WithField("job_uuid", job.UUID).
		WithField("status", nextStatus).
		Info("job updated successfully")
	return parsers.NewJobEntityFromModels(jobModel), nil
}

func (uc *updateJobUseCase) updateJob(ctx context.Context, jobModel *models.Job, jobLogModel *models.Log) error {
	logger := log.WithContext(ctx).WithField("job_uuid", jobModel.UUID)

	// Does current job belong to a parent/children chains?
	var parentJobUUID string
	if jobModel.InternalData.ParentJobUUID != "" {
		parentJobUUID = jobModel.InternalData.ParentJobUUID
	} else if jobModel.InternalData.RetryInterval != 0 {
		parentJobUUID = jobModel.UUID
	}

	prevLogModel := jobModel.Logs[len(jobModel.Logs)-1]
	err := database.ExecuteInDBTx(uc.db, func(tx database.Tx) error {
		// We should lock ONLY when there is children jobs
		if parentJobUUID != "" {
			logger.WithField("parent_job", parentJobUUID).Debug("lock parent job row for update")
			if err := tx.(store.Tx).Job().LockOneByUUID(ctx, parentJobUUID); err != nil {
				return err
			}

			// Refresh jobModel after lock to ensure nothing was updated
			refreshedJobModel, err := uc.db.Job().FindOneByUUID(ctx, jobModel.UUID, []string{}, false)
			if err != nil {
				return err
			}

			if refreshedJobModel.UpdatedAt != jobModel.UpdatedAt {
				errMessage := "job status was updated since user request was sent"
				logger.WithField("status", jobModel.Status).Error(errMessage)
				return errors.InvalidStateError(errMessage).ExtendComponent(updateJobComponent)
			}
		}

		if jobLogModel != nil {
			if err := tx.(store.Tx).Log().Insert(ctx, jobLogModel); err != nil {
				return err
			}

			jobModel.Logs = append(jobModel.Logs, jobLogModel)
			if updateNextJobStatus(prevLogModel.Status, jobLogModel.Status) {
				jobModel.Status = jobLogModel.Status
			}
		}

		if err := tx.(store.Tx).Job().Update(ctx, jobModel); err != nil {
			return err
		}

		// if we updated to MINED, we need to update the children and sibling jobs to NEVER_MINED
		if parentJobUUID != "" && jobLogModel != nil && jobLogModel.Status == utils.StatusMined {
			der := uc.updateChildrenUseCase.
				WithDBTransaction(tx.(store.Tx)).
				Execute(ctx, jobModel.UUID, parentJobUUID, utils.StatusNeverMined, []string{})
			if der != nil {
				return der
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func updateNextJobStatus(prevStatus, nextStatus string) bool {
	if nextStatus == utils.StatusResending {
		return false
	}
	if nextStatus == utils.StatusWarning {
		return false
	}
	if nextStatus == utils.StatusFailed && prevStatus == utils.StatusResending {
		return false
	}

	return true
}

func canUpdateStatus(nextStatus, status string) bool {
	switch nextStatus {
	case utils.StatusCreated:
		return false
	case utils.StatusStarted:
		return status == utils.StatusCreated
	case utils.StatusPending:
		return status == utils.StatusStarted || status == utils.StatusRecovering
	case utils.StatusResending:
		return status == utils.StatusPending || status == utils.StatusRecovering
	case utils.StatusRecovering:
		return status == utils.StatusStarted || status == utils.StatusPending || status == utils.StatusRecovering
	case utils.StatusMined, utils.StatusStored, utils.StatusNeverMined:
		return status == utils.StatusPending
	case utils.StatusFailed:
		return status == utils.StatusStarted || status == utils.StatusRecovering || status == utils.StatusPending
	default: // For warning, they can be added at any time
		return true
	}
}
