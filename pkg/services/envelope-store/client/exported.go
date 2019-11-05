package client

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/tracing/opentracing/jaeger"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/opentracing"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	grpcclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/grpc/client"
	evlpstore "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/services/envelope-store"
)

const component = "envelope-store.client"

type serviceName string

var (
	client   evlpstore.EnvelopeStoreClient
	conn     *grpc.ClientConn
	initOnce = &sync.Once{}
)

func Init(ctx context.Context) {
	initOnce.Do(func() {
		if client != nil {
			return
		}

		ctxWithValue := context.WithValue(ctx, serviceName("service-name"), viper.GetString(jaeger.ServiceNameViperKey))
		opentracing.Init(ctxWithValue)
		var err error
		conn, err = grpcclient.DialContextWithDefaultOptions(
			ctx,
			viper.GetString(EnvelopeStoreURLViperKey),
		)
		if err != nil {
			e := errors.FromError(err).ExtendComponent(component)
			log.WithError(e).Fatalf("%s: failed to dial grpc server", component)
		}

		client = evlpstore.NewEnvelopeStoreClient(conn)

		log.WithFields(log.Fields{
			"url": viper.GetString(EnvelopeStoreURLViperKey),
		}).Infof("%s: client ready", component)
	})
}

func Close() {
	_ = conn.Close()
}

func GlobalEnvelopeStoreClient() evlpstore.EnvelopeStoreClient {
	return client
}

// SetGlobalConfig sets Sarama global configuration
func SetGlobalEnvelopeStoreClient(c evlpstore.EnvelopeStoreClient) {
	client = c
}
