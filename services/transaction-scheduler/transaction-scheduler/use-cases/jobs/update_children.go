package jobs

import (
	"context"
	"fmt"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"
)

const updateChildrenComponent = "use-cases.update-children"

// createJobUseCase is a use case to create a new transaction job
type updateChildrenUseCase struct {
	db store.DB
}

// NewUpdateChildrenUseCase creates a new UpdateChildrenUseCase
func NewUpdateChildrenUseCase(db store.DB) usecases.UpdateChildrenUseCase {
	return &updateChildrenUseCase{
		db: db,
	}
}

func (uc updateChildrenUseCase) WithDBTransaction(dbtx store.Tx) usecases.UpdateChildrenUseCase {
	uc.db = dbtx
	return &uc
}

// Execute updates all children of a job to NEVER_MINED
func (uc *updateChildrenUseCase) Execute(ctx context.Context, jobUUID, parentJobUUID, nextStatus string, tenants []string) error {
	logger := log.WithContext(ctx).WithField("job_uuid", jobUUID).WithField("parent_job_uuid", parentJobUUID).
		WithField("tenants", tenants).WithField("next_status", nextStatus)
	logger.Debug("updating sibling and/or parent jobs")

	if !utils.IsFinalStatus(nextStatus) {
		errMsg := "expected final job status"
		err := errors.InvalidParameterError(errMsg).ExtendComponent(updateChildrenComponent)
		logger.WithError(err).Error("failed to update children jobs")
		return err
	}

	jobsToUpdate, err := uc.db.Job().Search(ctx, &entities.JobFilters{
		ParentJobUUID: parentJobUUID,
		Status:        utils.StatusPending,
	}, tenants)

	if err != nil {
		return errors.FromError(err).ExtendComponent(updateChildrenComponent)
	}

	for _, jobModel := range jobsToUpdate {
		// Skip mined job which trigger the update of sibling/children
		if jobModel.UUID == jobUUID {
			continue
		}

		jobModel.Status = nextStatus
		if err := uc.db.Job().Update(ctx, jobModel); err != nil {
			return errors.FromError(err).ExtendComponent(updateChildrenComponent)
		}

		jobLogModel := &models.Log{
			JobID:   &jobModel.ID,
			Status:  nextStatus,
			Message: fmt.Sprintf("sibling (or parent) job %s was mined instead", jobUUID),
		}

		if err := uc.db.Log().Insert(ctx, jobLogModel); err != nil {
			return errors.FromError(err).ExtendComponent(updateChildrenComponent)
		}

		logger.WithField("job", jobModel.UUID).
			WithField("status", nextStatus).Debug("updated children/sibling job successfully")
	}

	logger.Info("children (and/or parent) jobs updated successfully")
	return nil
}
