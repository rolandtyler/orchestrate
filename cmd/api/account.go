package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/ConsenSys/orchestrate/pkg/toolkit/database/postgres"
	kmclient "github.com/ConsenSys/orchestrate/services/key-manager/client"
	"github.com/go-pg/pg/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newAccountCmd() *cobra.Command {
	var db *pg.DB

	accountCmd := &cobra.Command{
		Use:   "account",
		Short: "Account management",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Set database connection
			opts, err := postgres.NewConfig(viper.GetViper()).PGOptions()
			if err != nil {
				return err
			}
			db = pg.Connect(opts)

			// Init QKM client
			kmclient.Init()
			return nil
		},
	}

	// Postgres flags
	postgres.PGFlags(accountCmd.Flags())
	kmclient.Flags(accountCmd.Flags())

	importCmd := &cobra.Command{
		Use:   "import",
		Short: "import accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return importAccounts(cmd.Context(), db, kmclient.GlobalClient())
		},
	}
	accountCmd.AddCommand(importCmd)

	return accountCmd
}

func importAccounts(ctx context.Context, db *pg.DB, client kmclient.KeyManagerClient) error {
	log.Debug("Loading accounts from Vault...")

	namespaces, err := client.ETHListNamespaces(ctx)
	if err != nil {
		log.WithError(err).Errorf("could not get list of namespaces")
		return err
	}

	var queryInsertItems []string
	for _, namespace := range namespaces {
		accounts, err2 := client.ETHListAccounts(ctx, namespace)
		if err2 != nil {
			log.WithField("namespace", namespace).WithError(err2).Errorf("Could not get list of accounts")
			return err2
		}

		for _, addr := range accounts {
			acc, err2 := client.ETHGetAccount(ctx, addr, namespace)
			if err2 != nil {
				log.WithField("namespace", namespace).WithField("address", addr).
					WithError(err2).Error("Could not get account")
				return err2
			}

			queryInsertItems = append(queryInsertItems, fmt.Sprintf("('%s', '%s', '%s', '%s', '{\"source\": \"kv-v2\"}')",
				acc.Namespace,
				acc.Address,
				acc.PublicKey,
				acc.CompressedPublicKey,
			))
		}
	}

	if len(queryInsertItems) > 0 {
		_, err = db.Exec("INSERT INTO accounts (tenant_id, address, public_key, compressed_public_key, attributes) VALUES " +
			strings.Join(queryInsertItems, ", ") + " on conflict do nothing")
		if err != nil {
			log.WithError(err).Error("Could not import accounts")
			return err
		}
	}

	log.WithField("accounts", len(queryInsertItems)).Info("accounts imported successfully")
	return nil
}
