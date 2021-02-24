package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration06(db migrations.DB) error {
	log.Debug("Applying migration 06...")
	_, err := db.Exec(`
CREATE INDEX transactions_hash_idx on transactions (hash);
CREATE INDEX schedules_tenant_id_uuid_idx on schedules (tenant_id, uuid);
CREATE INDEX logs_job_id_status_idx on logs (job_id, status);

ALTER TABLE jobs 
	ADD COLUMN is_parent BOOLEAN NOT NULL DEFAULT FALSE,
	ADD COLUMN status job_status;

UPDATE jobs j1
	SET is_parent=j1.internal_data->'parentJobUUID' is null;

UPDATE jobs j1
	SET status=l1.status
	FROM logs l1
	WHERE l1.id=(select MAX(id) from logs l2 where l2.job_id = j1.id and l2.status NOT IN ('WARNING', 'RECOVERING', 'RESENDING'));

CREATE INDEX jobs_parent_updated_at_idx on jobs (is_parent, updated_at);
CREATE INDEX jobs_chain_uuid_status_idx on jobs (chain_uuid, status);
CREATE INDEX jobs_schedule_id_idx on jobs (schedule_id);
`)
	if err != nil {
		log.WithError(err).Error("Could not apply migration")
		return err
	}
	log.Info("Migration 06 completed")

	return nil
}

func downMigration06(db migrations.DB) error {
	log.Debug("Rollback migration 06...")
	_, err := db.Exec(`
DROP INDEX transactions_hash_idx;
DROP INDEX schedules_tenant_id_uuid_idx;
DROP INDEX jobs_parent_updated_at_idx;
DROP INDEX jobs_chain_uuid_status_idx;

ALTER TABLE jobs 
	DROP COLUMN is_parent, DROP COLUMN status;

DROP INDEX logs_job_id_status_idx;
DROP INDEX jobs_schedule_id_idx;
`)
	if err != nil {
		log.WithError(err).Error("Could not apply rollback")
		return err
	}

	log.Info("Rollback 06 completed")
	return nil
}

func init() {
	Collection.MustRegisterTx(upMigration06, downMigration06)
}
