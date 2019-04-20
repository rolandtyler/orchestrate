package maxbalance

import (
	"context"
	"math/big"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/infra/ethereum.git/ethclient"
	"gitlab.com/ConsenSys/client/fr/core-stack/infra/faucet.git/faucet"
)

var (
	ctrl     *Controller
	config   *Config
	initOnce = sync.Once{}
)

// Init initialize BlackList Controller
func Init(ctx context.Context) {
	initOnce.Do(func() {
		// Set config if not yet set
		if config == nil {
			InitConfig(ctx)
		}

		// Initialize controller
		ctrl = NewController(config)

		log.WithFields(log.Fields{
			"controller":     "max-balance",
			"controller.max": ctrl.conf.MaxBalance.Text(10),
		}).Info("faucet: controller ready")
	})
}

// InitConfig initialize configuration
func InitConfig(ctx context.Context) {
	max, ok := big.NewInt(0).SetString(viper.GetString(faucetMaxViperKey), 10)
	if !ok {
		log.Fatalf("max-balance: invalid maximum balance %q", viper.GetString(faucetMaxViperKey))
	}

	// Initialize global MultiEthClient
	ethclient.Init(ctx)

	config = &Config{
		MaxBalance: max,
		BalanceAt:  ethclient.GlobalMultiClient().BalanceAt,
	}
}

// SetGlobalConfig sets global configuration
func SetGlobalConfig(c *Config) {
	config = c
}

// GlobalConfig returns global configuration
func GlobalConfig() *Config {
	return config
}

// GlobalController returns global blacklist controller
func GlobalController() *Controller {
	return ctrl
}

// SetGlobalController sets global blacklist controller
func SetGlobalController(controller *Controller) {
	initOnce.Do(func() {
		ctrl = controller
	})
}

// Control allows to control a CreditFunc with global MaxBalance
func Control(f faucet.CreditFunc) faucet.CreditFunc {
	return ctrl.Control(f)
}
