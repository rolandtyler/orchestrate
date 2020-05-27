package txcrafter

import (
	"context"
	"os"

	"github.com/containous/traefik/v2/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	chnregclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/client"
	registryclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/client"
	txcrafter "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-crafter"
)

var cmdErr error

func newRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run application",
		Run:   run,
		PreRun: func(cmd *cobra.Command, args []string) {
			utils.PreRunBindFlags(viper.GetViper(), cmd.Flags(), "tx-crafter")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			if err := errors.CombineErrors(cmdErr, cmd.Context().Err()); err != nil {
				os.Exit(1)
			}
		},
	}

	// Register Kafka flags
	broker.InitKafkaFlags(runCmd.Flags())
	broker.KafkaTopicTxCrafter(runCmd.Flags())
	broker.KafkaTopicTxRecover(runCmd.Flags())

	// Chain Registry
	chnregclient.Flags(runCmd.Flags())

	// Contract Registry
	registryclient.ContractRegistryURL(runCmd.Flags())

	return runCmd
}

func run(cmd *cobra.Command, _ []string) {
	ctx, cancel := context.WithCancel(cmd.Context())
	logger := log.FromContext(ctx)

	// Start microservice
	go func() {
		if err := <-txcrafter.Start(ctx); err != nil {
			cmdErr = errors.CombineErrors(cmdErr, err)
			logger.WithError(err).Errorf("Microservice raised an error")
		}
		cancel()
	}()

	// Process signals
	sig := utils.NewSignalListener(func(signal os.Signal) {
		cancel()
	})

	// Stop when get context canceled
	<-ctx.Done()
	err := txcrafter.Stop(ctx)
	if err != nil {
		cmdErr = errors.CombineErrors(cmdErr, err)
		logger.WithError(err).Errorf("Microservice did not shutdown properly")
	} else {
		logger.Infof("Microservice gracefully closed")
	}

	sig.Close()
}
