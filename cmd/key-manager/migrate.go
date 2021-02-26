package keymanager

import (
	kvv2 "github.com/ConsenSys/orchestrate/pkg/hashicorp/kv-v2"
	keymanager "github.com/ConsenSys/orchestrate/services/key-manager"
	"github.com/ConsenSys/orchestrate/services/key-manager/store"
	migrations2 "github.com/ConsenSys/orchestrate/services/key-manager/store/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newMigrateCmd create migrate command
func newMigrateCmd() *cobra.Command {
	var vault store.Vault
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migration of Vault secrets",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfg := keymanager.NewConfig(viper.GetViper())
			var err error
			vault, err = store.Build(cmd.Context(), cfg.Store)
			if err != nil {
				return err
			}
			return vault.HealthCheck()
		},
	}

	// Register Init command
	importSecretCmd := &cobra.Command{
		Use:   "import-secrets",
		Short: "Import secrets store in old Hashicorp vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Initialize v2 client
			cfg := kvv2.ConfigFromViper()
			v2Client, err := kvv2.NewClient(cfg.Config, cfg.SecretPath)
			if err != nil {
				return err
			}

			return migrations2.Kvv2ImportSecrets(cmd.Context(), vault, v2Client)
		},
	}

	kvv2.InitFlags(importSecretCmd.Flags())
	migrateCmd.AddCommand(importSecretCmd)

	return migrateCmd
}
