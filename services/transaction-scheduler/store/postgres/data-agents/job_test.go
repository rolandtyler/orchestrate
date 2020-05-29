// +build unit
// +build !race
// +build !integration

package dataagents

import (
	"context"
	"testing"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	pgTestUtils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database/postgres/testutils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models/testutils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/postgres/migrations"
)

type jobTestSuite struct {
	suite.Suite
	agents *PGAgents
	pg     *pgTestUtils.PGTestHelper
}

func TestPGJob(t *testing.T) {
	s := new(jobTestSuite)
	suite.Run(t, s)
}

func (s *jobTestSuite) SetupSuite() {
	s.pg , _ = pgTestUtils.NewPGTestHelper(nil, migrations.Collection)
	s.pg.InitTestDB(s.T())
}

func (s *jobTestSuite) SetupTest() {
	s.pg.UpgradeTestDB(s.T())
	s.agents = New(s.pg.DB)
}

func (s *jobTestSuite) TearDownTest() {
	s.pg.DowngradeTestDB(s.T())
}

func (s *jobTestSuite) TearDownSuite() {
	s.pg.DropTestDB(s.T())
}

func (s *jobTestSuite) TestPGJob_Insert() {
	ctx := context.Background()

	s.T().Run("should insert model successfully", func(t *testing.T) {
		job := testutils.FakeJob(0)
		err := insertJob(ctx, s.agents, job)
		assert.Nil(s.T(), err)

		assert.Nil(t, err)
		assert.NotEmpty(t, job.ID)
		assert.NotEmpty(t, job.Transaction.ID)
		assert.NotEmpty(t, job.Schedule.ID)
	})
	
	s.T().Run("should insert model without UUID successfully", func(t *testing.T) {
		job := testutils.FakeJob(0)
		job.UUID = ""
		err := insertJob(ctx, s.agents, job)
		assert.Nil(s.T(), err)

		assert.Nil(t, err)
		assert.NotEmpty(t, job.ID)
		assert.NotEmpty(t, job.Transaction.ID)
		assert.NotEmpty(t, job.Schedule.ID)
	})

	s.T().Run("should update model successfully", func(t *testing.T) {
		job := testutils.FakeJob(0)
		err := insertJob(ctx, s.agents, job)
		assert.Nil(s.T(), err)

		assert.Nil(t, err)
		assert.NotEmpty(t, job.ID)
		assert.NotEmpty(t, job.Transaction.ID)
		assert.NotEmpty(t, job.Schedule.ID)
	})
}

func (s *jobTestSuite) TestPGJob_Update() {
	ctx := context.Background()
	job := testutils.FakeJob(0)
	err := insertJob(ctx, s.agents, job)
	assert.Nil(s.T(), err)

	s.T().Run("should update model successfully", func(t *testing.T) {
		newTx := testutils.FakeTransaction()
		newSchedule := testutils.FakeSchedule("_")
		err := s.agents.Transaction().Insert(ctx, newTx)
		assert.Nil(t, err)
		err = s.agents.Schedule().Insert(ctx, newSchedule)
		assert.Nil(t, err)

		job.ScheduleID = &newSchedule.ID
		job.TransactionID = &newTx.ID
		err = s.agents.Job().Update(ctx, job)
		assert.Nil(t, err)
		assert.Equal(t, *job.TransactionID, newTx.ID)
		assert.Equal(t, *job.ScheduleID, newSchedule.ID)
	})

	s.T().Run("should fail to update job with missing ID", func(t *testing.T) {
		job.ID = 0
		err = s.agents.Job().Update(ctx, job)
		assert.True(t, errors.IsInvalidArgError(err))
	})
}

func (s *jobTestSuite) TestPGJob_FindOneByUUID() {
	ctx := context.Background()
	tenantID := "tenantID"
	job := testutils.FakeJob(0)
	job.Schedule.TenantID = tenantID
	err := insertJob(ctx, s.agents, job)
	assert.Nil(s.T(), err)

	s.T().Run("should get model successfully as tenant", func(t *testing.T) {
		jobRetrieved, err := s.agents.Job().FindOneByUUID(ctx, job.UUID, "_")

		assert.Nil(t, err)
		assert.NotEmpty(t, jobRetrieved.ID)
		assert.Equal(t, job.UUID, jobRetrieved.UUID)
		assert.Equal(t, job.Transaction.UUID, jobRetrieved.Transaction.UUID)
		assert.NotEmpty(t, jobRetrieved.Transaction.ID)
		assert.Equal(t, job.Logs[0].UUID, jobRetrieved.Logs[0].UUID)
		assert.NotEmpty(t, jobRetrieved.Logs[0].ID)
		assert.Equal(t, job.Schedule.UUID, jobRetrieved.Schedule.UUID)
		assert.Equal(t, job.Schedule.TenantID, jobRetrieved.Schedule.TenantID)
		assert.NotEmpty(t, jobRetrieved.Schedule.ID)
	})

	s.T().Run("should get model successfully as tenant", func(t *testing.T) {
		jobRetrieved, err := s.agents.Job().FindOneByUUID(ctx, job.UUID, tenantID)

		assert.Nil(t, err)
		assert.NotEmpty(t, jobRetrieved.ID)
	})

	s.T().Run("should get model successfully as admin", func(t *testing.T) {
		jobRetrieved, err := s.agents.Job().FindOneByUUID(ctx, job.UUID, multitenancy.DefaultTenantIDName)

		assert.Nil(t, err)
		assert.NotEmpty(t, jobRetrieved.ID)
		assert.Equal(t, job.UUID, jobRetrieved.UUID)
	})

	s.T().Run("should return NotFoundError if select fails", func(t *testing.T) {
		_, err := s.agents.Job().FindOneByUUID(ctx, "b6fe7a2a-1a4d-49ca-99d8-8a34aa495ef0", tenantID)
		assert.True(t, errors.IsNotFoundError(err))
	})
}

func (s *jobTestSuite) TestPGJob_ConnectionErr() {
	ctx := context.Background()

	// We drop the DB to make the test fail
	s.pg.DropTestDB(s.T())
	job := testutils.FakeJob(0)
	s.T().Run("should return PostgresConnectionError if insert fails", func(t *testing.T) {
		err := s.agents.Job().Insert(ctx, job)
		assert.True(t, errors.IsPostgresConnectionError(err))
	})

	s.T().Run("should return PostgresConnectionError if update fails", func(t *testing.T) {
		job.ID = 1
		err := s.agents.Job().Update(ctx, job)
		assert.True(t, errors.IsPostgresConnectionError(err))
	})

	s.T().Run("should return PostgresConnectionError if update fails", func(t *testing.T) {
		_, err := s.agents.Job().FindOneByUUID(ctx, job.UUID, job.Schedule.TenantID)
		assert.True(t, errors.IsPostgresConnectionError(err))
	})

	// We bring it back up
	s.pg.InitTestDB(s.T())
}

/**
Persist Job entity and its related entities
*/
func insertJob(ctx context.Context, agents *PGAgents, job *models.Job) error {
	if job.Schedule != nil {
		if err := agents.Schedule().Insert(ctx, job.Schedule); err != nil {
			return err
		}
	}

	if job.Transaction != nil {
		if err := agents.Transaction().Insert(ctx, job.Transaction); err != nil {
			return err
		}
	}

	if err := agents.Job().Insert(ctx, job); err != nil {
		return err
	}

	for idx := range job.Logs {
		job.Logs[idx].JobID = &job.ID
		if err := agents.Log().Insert(ctx, job.Logs[idx]); err != nil {
			return err
		}
	}

	return nil
}