package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration05(db migrations.DB) error {
	log.Debug("Applying migration 05...")
	_, err := db.Exec(`
CREATE TYPE job_status 
	AS ENUM ('CREATED', 'STARTED', 'PENDING', 'MINED', 'NEVER_MINED', 'RESENDING', 'STORED', 'RECOVERING', 'WARNING', 'FAILED');

CREATE TYPE job_type 
	AS ENUM ('eth://ethereum/transaction', 'eth://ethereum/rawTransaction', 'eth://orion/markingTransaction', 
	'eth://orion/eeaTransaction', 'eth://tessera/markingTransaction', 'eth://tessera/privateTransaction');

ALTER TABLE jobs
	ALTER COLUMN type TYPE job_type using type::job_type;

ALTER TABLE logs
	ALTER COLUMN status TYPE job_status using status::job_status;
`)
	if err != nil {
		log.WithError(err).Error("Could not apply migration")
		return err
	}
	log.Info("Migration 05 completed")

	return nil
}

func downMigration05(db migrations.DB) error {
	log.Debug("Rollback migration 05...")
	_, err := db.Exec(`
ALTER TABLE logs
	ALTER COLUMN status TYPE TEXT;

ALTER TABLE jobs
	ALTER COLUMN type TYPE TEXT;

DROP TYPE job_status;
DROP TYPE job_type;
`)
	if err != nil {
		log.WithError(err).Error("Could not apply rollback")
		return err
	}

	log.Info("Rollback 05 completed")
	return nil
}

func init() {
	Collection.MustRegisterTx(upMigration05, downMigration05)
}
