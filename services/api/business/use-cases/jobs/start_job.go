package jobs

import (
	"context"
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/database"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils/envelope"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/metrics"

	"github.com/Shopify/sarama"
	pkgsarama "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/log"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/parsers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models"
)

const startJobComponent = "use-cases.start-job"

// startJobUseCase is a use case to start a transaction job
type startJobUseCase struct {
	db            store.DB
	kafkaProducer sarama.SyncProducer
	topicsCfg     *pkgsarama.KafkaTopicConfig
	metrics       metrics.TransactionSchedulerMetrics
	logger        *log.Logger
}

// NewStartJobUseCase creates a new StartJobUseCase
func NewStartJobUseCase(
	db store.DB,
	kafkaProducer sarama.SyncProducer,
	topicsCfg *pkgsarama.KafkaTopicConfig,
	m metrics.TransactionSchedulerMetrics,
) usecases.StartJobUseCase {
	return &startJobUseCase{
		db:            db,
		kafkaProducer: kafkaProducer,
		topicsCfg:     topicsCfg,
		metrics:       m,
		logger:        log.NewLogger().SetComponent(startJobComponent),
	}
}

// Execute sends a job to the Kafka topic
func (uc *startJobUseCase) Execute(ctx context.Context, jobUUID string, tenants []string) error {
	logger := uc.logger.WithContext(ctx).WithField("job", jobUUID)
	logger.Debug("starting job")

	jobModel, err := uc.db.Job().FindOneByUUID(ctx, jobUUID, tenants, false)
	if err != nil {
		return errors.FromError(err).ExtendComponent(startJobComponent)
	}

	jobEntity := parsers.NewJobEntityFromModels(jobModel)
	if !canUpdateStatus(entities.StatusStarted, jobEntity.Status) {
		errMessage := "cannot start job at the current status"
		logger.WithField("status", jobEntity.Status).WithField("next_status", entities.StatusStarted).Error(errMessage)
		return errors.InvalidStateError(errMessage)
	}

	err = uc.updateStatus(ctx, jobModel, entities.StatusStarted, "")
	if err != nil {
		return errors.FromError(err).ExtendComponent(startJobComponent)
	}

	partition, offset, err := envelope.SendJobMessage(jobEntity, uc.kafkaProducer, uc.topicsCfg.Sender)
	if err != nil {
		errMsg := "failed to send job message"
		_ = uc.updateStatus(ctx, jobModel, entities.StatusFailed, errMsg)
		logger.WithError(err).Error(errMsg)
		return errors.FromError(err).ExtendComponent(startJobComponent)
	}

	logger.WithField("partition", partition).WithField("offset", offset).Info("job started successfully")

	return nil
}

func (uc *startJobUseCase) updateStatus(ctx context.Context, job *models.Job, status entities.JobStatus, msg string) error {
	job.Status = status
	jobLog := &models.Log{
		JobID:   &job.ID,
		Status:  status,
		Message: msg,
	}

	prevUpdatedAt := job.UpdatedAt
	prevStatus := job.Status
	err := database.ExecuteInDBTx(uc.db, func(tx database.Tx) error {
		if err := tx.(store.Tx).Job().Update(ctx, job); err != nil {
			return err
		}

		if err := tx.(store.Tx).Log().Insert(ctx, jobLog); err != nil {
			return errors.FromError(err).ExtendComponent(startJobComponent)
		}

		return nil
	})

	if err != nil {
		return err
	}

	uc.addMetrics(job.UpdatedAt.Sub(prevUpdatedAt), prevStatus, status, job.ChainUUID)
	return nil
}

func (uc *startJobUseCase) addMetrics(elapseTime time.Duration, previousStatus, nextStatus entities.JobStatus, chainUUID string) {
	baseLabels := []string{
		"chain_uuid", chainUUID,
	}

	uc.metrics.JobsLatencyHistogram().With(append(baseLabels,
		"prev_status", string(previousStatus),
		"status", string(nextStatus),
	)...).Observe(elapseTime.Seconds())
}
