package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration07(db migrations.DB) error {
	log.Debug("Applying migration 07...")
	_, err := db.Exec(`
ALTER TABLE jobs
	DROP CONSTRAINT jobs_schedule_id_fkey;

ALTER TABLE jobs
	DROP CONSTRAINT jobs_transaction_id_fkey;

ALTER TABLE logs
	DROP CONSTRAINT logs_job_id_fkey;

ALTER TABLE transaction_requests
	DROP CONSTRAINT transaction_requests_schedule_id_fkey;

DROP TRIGGER job_trigger ON jobs;
`)
	if err != nil {
		log.WithError(err).Error("Could not apply migration")
		return err
	}
	log.Info("Migration 07 completed")

	return nil
}

func downMigration07(db migrations.DB) error {
	log.Debug("Rollback migration 07...")
	_, err := db.Exec(`
CREATE TRIGGER job_trigger
	BEFORE UPDATE ON jobs
	FOR EACH ROW 
	EXECUTE PROCEDURE updated();

ALTER TABLE jobs
	ADD CONSTRAINT jobs_schedule_id_fkey FOREIGN KEY (schedule_id) REFERENCES schedules (id);

ALTER TABLE jobs
	ADD CONSTRAINT jobs_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES transactions (id);

ALTER TABLE logs
	ADD CONSTRAINT logs_job_id_fkey FOREIGN KEY (job_id) REFERENCES jobs (id);

ALTER TABLE transaction_requests
	ADD CONSTRAINT transaction_requests_schedule_id_fkey FOREIGN KEY (schedule_id) REFERENCES schedules (id);
`)
	if err != nil {
		log.WithError(err).Error("Could not apply rollback")
		return err
	}

	log.Info("Rollback 07 completed")
	return nil
}

func init() {
	Collection.MustRegisterTx(upMigration07, downMigration07)
}
