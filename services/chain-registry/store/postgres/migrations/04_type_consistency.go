package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration04(db migrations.DB) error {
	log.Debug("Applying migration 04...")
	_, err := db.Exec(`
ALTER TABLE chains
	ALTER COLUMN name TYPE TEXT,
	ALTER COLUMN tenant_id TYPE TEXT;
`)
	if err != nil {
		log.WithError(err).Error("Could not apply migration")
		return err
	}
	log.Info("Migration 04 completed")

	return nil
}

func downMigration04(db migrations.DB) error {
	log.Debug("Rollback migration 04...")
	_, err := db.Exec(`
ALTER TABLE chains
	ALTER COLUMN name TYPE VARCHAR(66),
	ALTER COLUMN tenant_id TYPE VARCHAR(66);
`)
	if err != nil {
		log.WithError(err).Error("Could not apply rollback")
		return err
	}

	log.Info("Rollback 04 completed")
	return nil
}

func init() {
	Collection.MustRegisterTx(upMigration04, downMigration04)
}
