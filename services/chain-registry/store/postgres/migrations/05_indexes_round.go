package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration05(db migrations.DB) error {
	log.Debug("Applying migration 05...")
	_, err := db.Exec(`
CREATE UNIQUE INDEX chain_uuid_idx ON chains (uuid);
CREATE INDEX private_tx_manager_chain_uuid_idx on private_tx_managers ("chain_uuid");
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
DROP INDEX private_tx_manager_chain_uuid_idx;
DROP INDEX chain_uuid_idx;
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
