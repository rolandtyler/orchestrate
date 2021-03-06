package dynamic

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	testutilsdynamic "github.com/consensys/orchestrate/pkg/toolkit/app/http/config/dynamic/testutils"
	mhandler "github.com/consensys/orchestrate/pkg/toolkit/app/http/handler/mock"
	mmiddleware "github.com/consensys/orchestrate/pkg/toolkit/app/http/middleware/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	traefikstatic "github.com/traefik/traefik/v2/pkg/config/static"
	traefiktypes "github.com/traefik/traefik/v2/pkg/types"
)

func TestBuilder(t *testing.T) {
	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	midBuilder := mmiddleware.NewMockBuilder(ctrlr)
	handlerBuilder := mhandler.NewMockBuilder(ctrlr)
	dashboardBuilder := mhandler.NewMockBuilder(ctrlr)
	accesslogBuilder := mmiddleware.NewMockBuilder(ctrlr)
	metricsBuilder := mmiddleware.NewMockBuilder(ctrlr)

	builder := NewBuilder(&traefikstatic.Configuration{
		HostResolver: &traefiktypes.HostResolverConfig{},
	}, nil)
	builder.Middleware = midBuilder
	builder.Handler = handlerBuilder
	builder.Metrics = metricsBuilder
	builder.accesslog = accesslogBuilder
	builder.dashboard = dashboardBuilder

	proxyHandler := mhandler.NewMockHandler(ctrlr)
	handlerBuilder.EXPECT().Build(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(proxyHandler, nil).Times(2)

	dashboardHandler := mhandler.NewMockHandler(ctrlr)
	dashboardBuilder.EXPECT().Build(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(dashboardHandler, nil).Times(1)

	midMockHandler := mhandler.NewMockHandler(ctrlr)
	mockMiddleware := mmiddleware.NewMockMiddleware(midMockHandler)

	midBuilder.EXPECT().Build(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockMiddleware, nil, nil).Times(3)

	accessLogMockHandler := mhandler.NewMockHandler(ctrlr)
	accessLogMiddleware := mmiddleware.NewMockMiddleware(accessLogMockHandler)
	accesslogBuilder.EXPECT().Build(gomock.Any(), gomock.Any(), gomock.Not(nil)).Return(accessLogMiddleware, nil, nil).Times(1)

	metricsMockHandler := mhandler.NewMockHandler(ctrlr)
	metricsMiddleware := mmiddleware.NewMockMiddleware(metricsMockHandler)
	metricsBuilder.EXPECT().Build(gomock.Any(), gomock.Any(), gomock.Any()).Return(metricsMiddleware, nil, nil).Times(3)

	routers, err := builder.Build(
		context.Background(),
		[]string{"ep-foo", "ep-bar"},
		testutilsdynamic.Config,
	)
	require.NoError(t, err, "Build shoud not error")

	// Call proxy
	midMockHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())
	proxyHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())
	metricsMockHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())

	req, _ := http.NewRequest(http.MethodGet, "http://proxy.com", nil)
	rw := httptest.NewRecorder()
	routers["ep-foo"].HTTP.ServeHTTP(rw, req)

	// Call dashboard
	accessLogMockHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Times(1)
	midMockHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Times(1)
	dashboardHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())
	metricsMockHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Times(1)

	req, _ = http.NewRequest(http.MethodGet, "http://dashboard.com", nil)
	rw = httptest.NewRecorder()
	routers["ep-foo"].HTTP.ServeHTTP(rw, req)

	req, _ = http.NewRequest(http.MethodGet, "http://unknown.com", nil)
	rw = httptest.NewRecorder()
	routers["ep-foo"].HTTP.ServeHTTP(rw, req)
}
