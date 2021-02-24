package migrations

import (
	"github.com/go-pg/migrations/v7"
	log "github.com/sirupsen/logrus"
)

func upMigration03(db migrations.DB) error {
	log.Debug("Applying migration 03...")
	_, err := db.Exec(`
ALTER TABLE artifacts
	ALTER COLUMN codehash TYPE TEXT;

ALTER TABLE codehashes
	ALTER COLUMN codehash TYPE TEXT,
	ALTER COLUMN address TYPE TEXT,
	ALTER COLUMN chain_id TYPE TEXT;

ALTER TABLE methods
	ALTER COLUMN codehash TYPE TEXT,
	ALTER COLUMN selector TYPE TEXT;

ALTER TABLE events
	ALTER COLUMN codehash TYPE TEXT,
	ALTER COLUMN sig_hash TYPE TEXT;

ALTER TABLE repositories
	ALTER COLUMN name TYPE TEXT;

ALTER TABLE tags
	ALTER COLUMN name TYPE TEXT;
`)
	if err != nil {
		log.WithError(err).Error("Could not apply migration")
		return err
	}
	log.Info("Migration 03 completed")

	return nil
}

func downMigration03(db migrations.DB) error {
	log.Debug("Rollback migration 03...")
	_, err := db.Exec(`
ALTER TABLE artifacts
	ALTER COLUMN codehash TYPE CHAR(66);

ALTER TABLE codehashes
	ALTER COLUMN codehash TYPE CHAR(66),
	ALTER COLUMN address TYPE char(42),
	ALTER COLUMN chain_id TYPE VARCHAR(66);

ALTER TABLE methods
	ALTER COLUMN codehash TYPE CHAR(66),
	ALTER COLUMN selector TYPE CHAR(10);

ALTER TABLE events
	ALTER COLUMN codehash TYPE CHAR(66),
	ALTER COLUMN sig_hash TYPE CHAR(66);

ALTER TABLE repositories
	ALTER COLUMN name TYPE VARCHAR(66);

ALTER TABLE tags
	ALTER COLUMN name TYPE VARCHAR(66);
`)
	if err != nil {
		log.WithError(err).Error("Could not apply rollback")
		return err
	}

	log.Info("Rollback 03 completed")
	return nil
}

func init() {
	Collection.MustRegisterTx(upMigration03, downMigration03)
}
