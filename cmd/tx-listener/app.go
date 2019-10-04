package txlistener

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/ethereum/ethclient"
	txlconfig "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/ethereum/tx-listener/handler/base"
	txlhandler "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/ethereum/tx-listener/handler/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/ethereum/tx-listener/listener"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/enricher"
	envelopeloader "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/envelope/loader"
	receiptloader "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/loader/receipt"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/logger"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/opentracing"
	producer "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/producer/tx-listener"
	injector "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/trace-injector"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/server/metrics"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/tracing/opentracing/jaeger"
)

var (
	app       = common.NewApp()
	startOnce = &sync.Once{}
)

type serviceName string

func initHandlers(ctx context.Context) {
	common.InParallel(
		// Initialize Jaeger
		func() {
			ctxWithValue := context.WithValue(ctx, serviceName("service-name"), viper.GetString(jaeger.ServiceNameViperKey))
			opentracing.Init(ctxWithValue)
		},
		// Initialize Jaeger trace injector
		func() {
			ctxWithValue := context.WithValue(ctx, serviceName("service-name"), viper.GetString(jaeger.ServiceNameViperKey))
			injector.Init(ctxWithValue)
		},
		// Initialize store
		func() {
			envelopeloader.Init(ctx)
		},
		// Initialize enricher
		func() {
			enricher.Init(ctx)
		},
		// Initialize producer
		func() {
			producer.Init(ctx)
		},
	)
}

func initComponents(ctx context.Context) {
	common.InParallel(
		func() {
			engine.Init(ctx)
		},
		func() {
			initHandlers(ctx)
		},
		func() {
			broker.InitSyncProducer(ctx)
		},
		func() {
			listener.Init(ctx)
		},
	)
}

func registerHandlers() {
	// Generic handlers on every worker
	engine.Register(logger.Logger)

	// Specific handlers to tx-listener
	engine.Register(opentracing.GlobalHandler())
	engine.Register(receiptloader.Loader)
	engine.Register(envelopeloader.GlobalHandler())
	engine.Register(opentracing.GlobalHandler())
	engine.Register(producer.GlobalHandler())
	engine.Register(injector.GlobalHandler())
}

// Start starts application
func Start(ctx context.Context) {
	startOnce.Do(func() {
		cancelCtx, cancel := context.WithCancel(ctx)
		go metrics.StartServer(ctx, cancel, app.IsAlive, app.IsReady)

		// Initialize all components of the server
		initComponents(cancelCtx)

		// Register all Handlers
		registerHandlers()

		// Indicate that application is ready
		// TODO: we need to update so SetReady can be called when Consume has finished to Setup
		app.SetReady(true)

		// Create handler
		conf, err := txlconfig.NewConfig()
		if err != nil {
			log.WithError(err).Fatalf("listener: could not load config")
		}
		h := txlhandler.NewHandler(engine.GlobalEngine(), broker.GlobalClient(), broker.GlobalSyncProducer(), conf)

		// Start Listening
		chains := ethclient.GlobalClient().Networks(cancelCtx)
		err = listener.Listen(cancelCtx, chains, h)
		if err != nil {
			log.WithError(err).Error("exiting loop with error")
		}
	})
}
