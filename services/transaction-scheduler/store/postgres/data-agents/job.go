package dataagents

import (
	"context"
	"fmt"

	pg "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database/postgres"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"
)

const jobDAComponent = "data-agents.job"

// PGJob is a job data agent for PostgreSQL
type PGJob struct {
	db pg.DB
}

// NewPGJob creates a new pgJob
func NewPGJob(db pg.DB) *PGJob {
	return &PGJob{db: db}
}

// Insert Inserts a new job in DB
func (agent *PGJob) Insert(ctx context.Context, job *models.Job) error {
	if job.UUID == "" {
		job.UUID = uuid.NewV4().String()
	}

	if job.Transaction != nil && job.TransactionID == nil {
		job.TransactionID = &job.Transaction.ID
	}
	if job.Schedule != nil && job.ScheduleID == nil {
		job.ScheduleID = &job.Transaction.ID
	}

	agent.db.ModelContext(ctx, job)
	err := pg.Insert(ctx, agent.db, job)
	if err != nil {
		return errors.FromError(err).ExtendComponent(jobDAComponent)
	}

	return nil
}

// Insert Inserts a new job in DB
func (agent *PGJob) Update(ctx context.Context, job *models.Job) error {
	if job.ID == 0 {
		return errors.InvalidArgError("cannot update job with missing ID")
	}

	if job.Transaction != nil && job.TransactionID == nil {
		job.TransactionID = &job.Transaction.ID
	}
	if job.Schedule != nil && job.ScheduleID == nil {
		job.ScheduleID = &job.Transaction.ID
	}

	agent.db.ModelContext(ctx, job)
	err := pg.Update(ctx, agent.db, job)
	if err != nil {
		return errors.FromError(err).ExtendComponent(jobDAComponent)
	}

	return nil
}

// FindOneByUUID gets a job by UUID
func (agent *PGJob) FindOneByUUID(ctx context.Context, jobUUID, tenantID string) (*models.Job, error) {
	job := &models.Job{}

	query := agent.db.ModelContext(ctx, job).
		Where("job.uuid = ?", jobUUID).
		Relation("Transaction").
		Relation("Schedule").
		Relation("Logs")

	if tenantID != multitenancy.DefaultTenantIDName {
		query.Where("schedule.tenant_id = ?", tenantID)
	}

	err := pg.Select(ctx, query)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(jobDAComponent)
	}

	return job, nil
}

func (agent *PGJob) Search(ctx context.Context, filters map[string]string, tenantID string) ([]*models.Job, error) {
	jobs := []*models.Job{}

	query := agent.db.ModelContext(ctx, &jobs).
		Relation("Transaction").
		Relation("Schedule").
		Relation("Logs")

	for attr, value := range filters {
		query = query.Where(fmt.Sprintf("job.%s = ?", attr), value)
	}

	if tenantID != multitenancy.DefaultTenantIDName {
		query.Where("schedule.tenant_id = ?", tenantID)
	}

	err := pg.Select(ctx, query)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(jobDAComponent)
	}

	return jobs, nil
}