package app

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/common"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/cmd_tmp/tx-crafter/handlers/loader"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/cmd_tmp/tx-crafter/handlers/logger"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/cmd_tmp/tx-crafter/handlers/offset"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/cmd_tmp/tx-crafter/handlers/opentracing"
	server "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/http/healthcheck"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-crafter.git/handlers"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-crafter.git/handlers/crafter"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-crafter.git/handlers/faucet"
	gasestimator "gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-crafter.git/handlers/gas-estimator"
	gaspricer "gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-crafter.git/handlers/gas-pricer"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-crafter.git/handlers/producer"
)

var (
	app       *App
	startOnce = &sync.Once{}
)

func init() {
	// Create app
	app = NewApp()

	// Set Kafka Group value
	viper.Set("kafka.group", "group-crafter")
}

func startServer(ctx context.Context) {
	// Initialize server
	server.Init(ctx)

	// Register Healthcheck
	server.Enhance(healthcheck.HealthCheck(app))

	// Start Listening
	_ = server.ListenAndServe()
}

func initConsumerGroup(ctx context.Context) {
	common.InParallel(
		// Initialize Engine
		func() {
			engine.Init(ctx)
		},
		// Initialize Handlers
		func() {
			handlers.Init(ctx)
		},
		// Initialize ConsumerGroup
		func() {
			broker.InitConsumerGroup(ctx)
		},
	)
	// Register handlers on engine
	// Generic handlers on every worker
	engine.Register(logger.Logger)
	engine.Register(loader.Loader)
	engine.Register(offset.Marker)
	engine.Register(producer.GlobalHandler())
	engine.Register(opentracing.GlobalHandler())

	// Specific handlers tk Tx-Crafter worker
	engine.Register(faucet.GlobalHandler())
	engine.Register(crafter.GlobalHandler())
	engine.Register(gaspricer.GlobalHandler())
	engine.Register(gasestimator.GlobalHandler())
}

// Start starts application
func Start(ctx context.Context) {
	startOnce.Do(func() {

		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			// Start Server
			startServer(ctx)
			cancel()
		}()

		// Initialize ConsumerGroup
		initConsumerGroup(cancelCtx)

		// Indicate that application is ready
		// TODO: we need to update so ready can append when Consume has finished to Setup
		app.ready.Store(true)

		// Start consuming on topic tx-crafter
		err := broker.Consume(
			cancelCtx,
			[]string{
				viper.GetString("kafka.topic.crafter"),
			},
			broker.NewEngineConsumerGroupHandler(engine.GlobalEngine()),
		)
		if err != nil {
			log.WithError(err).Error("worker: error on consumer")
		}
	})
}
