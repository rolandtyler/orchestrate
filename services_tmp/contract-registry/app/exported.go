package app

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	grpcserver "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/grpc/server"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/http/healthcheck"
	types "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/services/contract-registry"
	services "gitlab.com/ConsenSys/client/fr/core-stack/service/ethereum.git/abi/registry"
)

var (
	app       *App
	startOnce = &sync.Once{}
)

func init() {
	// Create app
	app = NewApp()
}

// Run application
func Start(ctx context.Context) {
	startOnce.Do(func() {
		// Initialize GRPC Server service
		services.Init(ctx)

		// Initialize GRPC server
		grpcserver.AddEnhancers(
			func(s *grpc.Server) *grpc.Server {
				types.RegisterRegistryServer(s, services.GlobalRegistry())
				return s
			},
		)
		grpcserver.Init(ctx)

		// Initialize HTTP server for healthchecks
		http.Init(ctx)
		http.Enhance(healthcheck.HealthCheck(app))

		// Indicate that application is ready
		app.ready.Store(true)

		// Start listening
		err := grpcserver.ListenAndServe()
		if err != nil {
			log.WithError(err).Error("app: error listening")
		}
	})
}

// Close gracefully stops the application
func Close(ctx context.Context) {
	log.Warn("app: stopping...")
	err := grpcserver.GracefulStop(ctx)
	if err != nil {
		log.WithError(err).Error("app: error stopping application")
	}
}
