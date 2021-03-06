package app

import (
	"context"
	"testing"
	"time"

	"github.com/consensys/orchestrate/pkg/toolkit/app/http/configwatcher"
	mockwatcher "github.com/consensys/orchestrate/pkg/toolkit/app/http/configwatcher/mock"
	mockhttp "github.com/consensys/orchestrate/pkg/toolkit/app/http/router/mock"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/toolkit/app/metrics/mock"
	"github.com/consensys/orchestrate/pkg/toolkit/app/metrics/registry"
	tcpmetrics "github.com/consensys/orchestrate/pkg/toolkit/tcp/metrics"
	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	traefikstatic "github.com/traefik/traefik/v2/pkg/config/static"
)

func newTestConfig() *Config {
	return &Config{
		HTTP: &HTTP{
			EntryPoints: map[string]*traefikstatic.EntryPoint{
				"test-ep": {
					Address: "127.0.0.1:1",
					Transport: &traefikstatic.EntryPointsTransport{
						RespondingTimeouts: &traefikstatic.RespondingTimeouts{},
						LifeCycle:          &traefikstatic.LifeCycle{},
					},
				},
			},
		},
		Watcher: &configwatcher.Config{
			ProvidersThrottleDuration: time.Millisecond,
		},
		Metrics: registry.NewConfig(viper.GetViper()),
	}
}

func TestApp(t *testing.T) {
	ctx := context.Background()
	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	httpBuilder := mockhttp.NewMockBuilder(ctrlr)

	watcher := mockwatcher.NewMockWatcher(ctrlr)

	reg := mock.NewMockRegistry(ctrlr)
	app := newApp(newTestConfig(), httpBuilder, watcher, reg, log.NewLogger())

	reg.EXPECT().Add(gomock.AssignableToTypeOf(tcpmetrics.NewTCPMetrics(nil)))
	watcher.EXPECT().AddListener(gomock.Any()).Times(2)
	watcher.EXPECT().Run(gomock.Any())

	err := app.Start(ctx)
	require.NoError(t, err, "App should have started properly")

	// Wait for application to properly start
	time.Sleep(100 * time.Millisecond)

	cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	err = app.Stop(cctx)
	assert.NoError(t, err, "App should have stop properly")

	watcher.EXPECT().Close()
	err = app.Close()
	assert.NoError(t, err, "App should have closed properly")
}
