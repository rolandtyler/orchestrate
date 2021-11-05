package e2e

import (
	"context"
	"github.com/consensys/orchestrate/pkg/sdk/client"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"net/http"
)

type Environment struct {
	Ctx    context.Context
	Logger *log.Logger
	Client client.OrchestrateClient
	Config *Config
}

func NewEnvironment() (*Environment, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	logger := log.NewLogger().SetComponent("e2e")

	orchestrateClient := client.NewHTTPClient(
		&http.Client{Transport: NewTestHttpTransport("", cfg.ApiKey)},
		client.NewConfig(cfg.ApiURL, nil),
	)

	return &Environment{
		Ctx:    context.Background(),
		Logger: logger,
		Client: orchestrateClient,
		Config: cfg,
	}, nil
}

func (e *Environment) SetupResources() error {
	return nil
}

func (e *Environment) TeardownResources() error {
	return nil
}
