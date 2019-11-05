package txsender

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/tracing/opentracing/jaeger"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/loader/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/logger"
	noncechecker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/nonce/checker"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/offset"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/opentracing"
	producer "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/producer/tx-sender"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/sender"
	injector "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/trace-injector"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
	server "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/http/healthcheck"
)

var (
	app       *common.App
	startOnce = &sync.Once{}
)

var l = createLogger("worker")

func init() {
	// Create app
	app = common.NewApp()
}

func startServer(ctx context.Context) {
	// Initialize server
	server.Init(ctx)

	// Register Healthcheck
	server.Enhance(healthcheck.HealthCheck(app))

	// Start Listening
	_ = server.ListenAndServe()
}

type serviceName string

func initHandlers(ctx context.Context) {
	common.InParallel(
		// Initialize Jaeger tracer
		func() {
			ctxWithValue := context.WithValue(ctx, serviceName("service-name"), viper.GetString(jaeger.ServiceNameViperKey))
			opentracing.Init(ctxWithValue)
		},
		// Initialize Jaeger tracer injector
		func() {
			ctxWithValue := context.WithValue(ctx, serviceName("service-name"), viper.GetString(jaeger.ServiceNameViperKey))
			injector.Init(ctxWithValue)
		},
		// Initialize sender
		func() { sender.Init(ctx) },
		// Initialize nonce manager
		func() { noncechecker.Init(ctx) },
		// Initialize producer
		func() { producer.Init(ctx) },
	)
}

func initComponents(ctx context.Context) {
	common.InParallel(
		// Initialize Engine
		func() { engine.Init(ctx) },

		// Initialize Handlers
		func() { initHandlers(ctx) },

		// Initialize ConsumerGroup
		func() {
			// Set Kafka Group value
			viper.Set(broker.KafkaGroupViperKey, "group-sender")
			broker.InitConsumerGroup(ctx)
		},
	)
}

func registerHandlers() {
	// Generic handlers on every worker
	engine.Register(opentracing.GlobalHandler())
	engine.Register(logger.Logger)
	engine.Register(sarama.Loader)
	engine.Register(offset.Marker)
	engine.Register(opentracing.GlobalHandler())
	engine.Register(producer.GlobalHandler())
	engine.Register(injector.GlobalHandler())
	engine.Register(noncechecker.GlobalRecoveryStatusSetter())

	// Specific handlers tk Sender worker
	engine.Register(noncechecker.GlobalChecker())
	engine.Register(sender.GlobalHandler())
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

		// Initialize all components of the server
		initComponents(cancelCtx)

		// Register all Handlers
		registerHandlers()

		// Indicate that application is ready
		// TODO: we need to update so ready can append when Consume has finished to Setup
		app.SetReady(true)

		topics := []string{
			viper.GetString(broker.TxSenderViperKey),
		}
		l.WithFields(log.Fields{
			"topics": topics,
		}).Info("connecting")
		// Start consuming on topic tx-sender
		err := broker.Consume(
			cancelCtx,
			topics,
			broker.NewEngineConsumerGroupHandler(engine.GlobalEngine()),
		)
		if err != nil {
			l.WithError(err).Error("error on consumer")
		}
	})
}

func createLogger(name string) *log.Entry {
	return log.WithField("name", name)
}
