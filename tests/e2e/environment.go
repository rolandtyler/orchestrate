package e2e

import (
	"context"
	"fmt"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/sdk/client"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/toolkit/ethclient/rpc"
	"github.com/consensys/orchestrate/pkg/types/api"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/consensys/orchestrate/tests/service/e2e/cucumber/alias"
	utils3 "github.com/consensys/orchestrate/tests/utils"
	"github.com/spf13/viper"
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

func (e *Environment) importTestAccounts() error {
	nodes := append(testData.Nodes.Besu, testData.Nodes.Quorum...)
	nodes = append(nodes, testData.Nodes.Geth...)

	for idx := range nodes {
		node := nodes[idx]
		for _, privKey := range node.FundedPrivateKeys {
			resp, err := orchestrateClient.ImportAccount(ctx, &api.ImportAccountRequest{
				PrivateKey: privKey,
			})

			if err != nil {
				if errors.IsAlreadyExistsError(err) || errors.IsConflictedError(err) || errors.IsDependencyFailureError(err) {
					logger.WithError(err).WithField("priv_key", privKey.String()).Warn("imported account is duplicated")
					continue
				}

				logger.WithError(err).WithField("priv_key", privKey.String()).Error("failed to import account")
				return err
			}

			logger.WithField("address", resp.Address).Info("account imported successfully")
		}
	}

	return nil
}

func (e *Environment) registerTestChains() error {
	aliases := alias.GlobalAliasRegistry()
	logger := log.FromContext(ctx)
	orchestrateClient := client.GlobalClient()
	ec := rpc.GlobalClient()
	proxyHost := viper.GetString(client.URLViperKey)

	reqs := map[string]*api.RegisterChainRequest{}
	for idx := range testData.Nodes.Besu {
		node := testData.Nodes.Besu[idx]
		reqs[fmt.Sprintf("besu%d", idx)] = &api.RegisterChainRequest{
			URLs: node.URLs,
			Name: fmt.Sprintf("besu-%s", utils.RandString(5)),
		}
	}

	for idx := range testData.Nodes.Geth {
		node := testData.Nodes.Geth[idx]
		reqs[fmt.Sprintf("geth%d", idx)] = &api.RegisterChainRequest{
			URLs: node.URLs,
			Name: fmt.Sprintf("geth-%s", utils.RandString(5)),
		}
	}

	for idx := range testData.Nodes.Quorum {
		node := testData.Nodes.Quorum[idx]
		if len(node.URLs) == 0 {
			continue
		}
		req := &api.RegisterChainRequest{
			URLs: node.URLs,
			Name: fmt.Sprintf("quorum-%s", utils.RandString(5)),
		}
		if node.PrivateTxManager.URL != "" {
			req.PrivateTxManager = &api.PrivateTxManagerRequest{
				URL:  node.PrivateTxManager.URL,
				Type: entities.TesseraChainType,
			}
		}
		reqs[fmt.Sprintf("quorum%d", idx)] = req
	}

	chainUUIDs := map[string]string{}
	for chainAlias, req := range reqs {
		resp, err := orchestrateClient.RegisterChain(ctx, req)
		if err != nil {
			logger.WithField("name", req.Name).WithError(err).Error("failed to register chain")
			return chainUUIDs, err
		}

		logger.WithField("name", req.Name).WithField("uuid", resp.UUID).WithField("alias", chainAlias).
			Info("chain registered successfully")
		chainUUIDs[req.Name] = resp.UUID

		aliases.Set(resp.UUID, fmt.Sprintf("chain.%s.UUID", chainAlias))
		aliases.Set(resp.Name, fmt.Sprintf("chain.%s.Name", chainAlias))
	}

	for _, chainUUID := range chainUUIDs {
		err := utils3.WaitForProxy(ctx, proxyHost, chainUUID, ec)
		if err != nil {
			logger.WithField("uuid", chainUUID).WithError(err).Error("failed to wait for proxy chain")
			return chainUUIDs, err
		}
	}

	return chainUUIDs, nil
}

func (e *Environment) removeTestChains() error {
	orchestrateClient := client.GlobalClient()
	logger := log.FromContext(ctx)
	for chainName, chainUUID := range chainUUIDs {
		err := orchestrateClient.DeleteChain(ctx, chainUUID)
		if err != nil {
			logger.WithField("uuid", chainUUID).WithField("name", chainName).
				WithError(err).Error("failed to remove test chain")
			return err
		}

		logger.WithField("uuid", chainUUID).WithField("name", chainName).
			Info("test chain was removed successfully")
	}

	return nil
}
