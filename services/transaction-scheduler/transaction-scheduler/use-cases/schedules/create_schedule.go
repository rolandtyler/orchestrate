package schedules

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/parsers"
)

//go:generate mockgen -source=create_schedule.go -destination=mocks/create_schedule.go -package=mocks

const createScheduleComponent = "use-cases.create-schedule"

type CreateScheduleUseCase interface {
	Execute(ctx context.Context, schedule *entities.Schedule, tenantID string) (*entities.Schedule, error)
	WithDBTransaction(dbtx store.Tx) CreateScheduleUseCase
}

// createScheduleUseCase is a use case to create a new transaction schedule
type createScheduleUseCase struct {
	db store.DB
}

// NewCreateScheduleUseCase creates a new CreateScheduleUseCase
func NewCreateScheduleUseCase(db store.DB) CreateScheduleUseCase {
	return &createScheduleUseCase{
		db: db,
	}
}

func (uc createScheduleUseCase) WithDBTransaction(dbtx store.Tx) CreateScheduleUseCase {
	uc.db = dbtx
	return &uc
}

// Execute validates and creates a new transaction schedule
func (uc *createScheduleUseCase) Execute(ctx context.Context, schedule *entities.Schedule, tenantID string) (*entities.Schedule, error) {
	log.WithContext(ctx).Debug("creating new schedule")

	scheduleModel := parsers.NewScheduleModelFromEntities(schedule, tenantID)

	if scheduleModel.TransactionRequest != nil && scheduleModel.TransactionRequest.IdempotencyKey != "" {
		txRequest, err := uc.db.TransactionRequest().
			FindOneByIdempotencyKey(ctx, scheduleModel.TransactionRequest.IdempotencyKey)

		if err != nil {
			return nil, errors.FromError(err).ExtendComponent(createScheduleComponent)
		}

		scheduleModel.TransactionRequestID = &txRequest.ID
	}

	if err := uc.db.Schedule().Insert(ctx, scheduleModel); err != nil {
		return nil, errors.FromError(err).ExtendComponent(createScheduleComponent)
	}

	log.WithContext(ctx).
		WithField("schedule_uuid", scheduleModel.UUID).
		Info("schedule created successfully")

	return parsers.NewScheduleEntityFromModels(scheduleModel), nil
}
