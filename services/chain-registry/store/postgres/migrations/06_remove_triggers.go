package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration06(db migrations.DB) error {
	log.Debug("Applying migration 06...")
	_, err := db.Exec(`
DROP TRIGGER faucet_trigger ON faucets;

DROP TRIGGER chain_trigger ON chains;
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
CREATE TRIGGER chain_trigger
	BEFORE UPDATE ON chains
	FOR EACH ROW 
	EXECUTE PROCEDURE updated();

CREATE TRIGGER faucet_trigger
	BEFORE UPDATE ON faucets
	FOR EACH ROW 
	EXECUTE PROCEDURE updated();
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
