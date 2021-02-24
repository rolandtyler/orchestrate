package jobs

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases"
	utils2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/utils"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	pkgsarama "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"
)

const startJobComponent = "use-cases.start-job"

// startJobUseCase is a use case to start a transaction job
type startJobUseCase struct {
	db            store.DB
	kafkaProducer sarama.SyncProducer
	topicsCfg     *pkgsarama.KafkaTopicConfig
}

// NewStartJobUseCase creates a new StartJobUseCase
func NewStartJobUseCase(db store.DB, kafkaProducer sarama.SyncProducer, topicsCfg *pkgsarama.KafkaTopicConfig) usecases.StartJobUseCase {
	return &startJobUseCase{
		db:            db,
		kafkaProducer: kafkaProducer,
		topicsCfg:     topicsCfg,
	}
}

// Execute sends a job to the Kafka topic
func (uc *startJobUseCase) Execute(ctx context.Context, jobUUID string, tenants []string) error {
	logger := log.WithContext(ctx).WithField("job_uuid", jobUUID)
	logger.Debug("starting job")

	jobModel, err := uc.db.Job().FindOneByUUID(ctx, jobUUID, tenants, false)
	if err != nil {
		return errors.FromError(err).ExtendComponent(startJobComponent)
	}

	if !canUpdateStatus(utils.StatusStarted, jobModel.Status) {
		errMessage := "cannot start job at the current status"
		logger.WithField("status", jobModel.Status).WithField("next_status", utils.StatusStarted).Error(errMessage)
		return errors.InvalidStateError(errMessage)
	}

	var msgTopic string
	switch {
	case jobModel.Type == utils.EthereumRawTransaction:
		msgTopic = uc.topicsCfg.Sender
	default:
		msgTopic = uc.topicsCfg.Crafter
	}

	err = uc.updateStatus(ctx, jobModel, utils.StatusStarted, "")
	if err != nil {
		return err
	}

	partition, offset, err := utils2.SendJobMessage(ctx, jobModel, uc.kafkaProducer, msgTopic)
	if err != nil {
		errMsg := "failed to send job message"
		_ = uc.updateStatus(ctx, jobModel, utils.StatusFailed, errMsg)
		logger.WithError(err).Error(errMsg)
		return errors.FromError(err).ExtendComponent(startJobComponent)
	}

	logger.WithField("partition", partition).WithField("offset", offset).Info("job started successfully")

	return nil
}

func (uc *startJobUseCase) updateStatus(ctx context.Context, job *models.Job, status, msg string) error {
	job.Status = status
	jobLog := &models.Log{
		JobID:   &job.ID,
		Status:  status,
		Message: msg,
	}

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

	return nil
}
