package chainregistry

import (
	"context"
	"time"

	orchestrateclient "github.com/consensys/orchestrate/pkg/sdk/client"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/types/api"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"

	"github.com/cenkalti/backoff/v4"
	"github.com/consensys/orchestrate/services/tx-listener/dynamic"
	"github.com/traefik/traefik/v2/pkg/job"
	"github.com/traefik/traefik/v2/pkg/safe"
)

const component = "tx-listener.chain-registry.provider"

type Provider struct {
	client orchestrateclient.OrchestrateClient
	conf   *Config
}

func (p *Provider) Run(ctx context.Context, configInput chan<- *dynamic.Message) error {
	return p.runWithRetry(
		ctx,
		configInput,
		backoff.WithContext(job.NewBackOff(backoff.NewExponentialBackOff()), ctx),
	)
}

func (p *Provider) runWithRetry(ctx context.Context, configInput chan<- *dynamic.Message, bckff backoff.BackOff) error {
	logger := log.FromContext(ctx).SetComponent(component)
	ctx = log.With(ctx, logger)
	return backoff.RetryNotify(
		safe.OperationWithRecover(func() error {
			return p.run(ctx, configInput)
		}),
		bckff,
		func(err error, d time.Duration) {
			log.FromContext(ctx).WithError(err).Warnf("provider restarting in... %v", d)
		},
	)
}

func (p *Provider) run(ctx context.Context, configInput chan<- *dynamic.Message) (err error) {
	ticker := time.NewTicker(p.conf.RefreshInterval)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ticker.C:
			var chains []*api.ChainResponse
			chains, err = p.client.SearchChains(ctx, &entities.ChainFilters{})
			if err != nil {
				log.FromContext(ctx).WithError(err).Error("failed to fetch chains from chain registry")
				break loop
			}
			configInput <- p.buildConfiguration(ctx, chains)
		case <-ctx.Done():
		}
	}

	return
}

func (p *Provider) buildConfiguration(ctx context.Context, chains []*api.ChainResponse) *dynamic.Message {
	msg := &dynamic.Message{
		Provider: "chain-registry",
		Configuration: &dynamic.Configuration{
			Chains: make(map[string]*dynamic.Chain),
		},
	}

	for _, chain := range chains {
		duration, err := time.ParseDuration(chain.ListenerBackOffDuration)
		if err != nil {
			log.FromContext(ctx).
				WithField("tenant_id", chain.TenantID).
				WithField("owner_id", chain.OwnerID).
				WithField("chain", chain.UUID).
				Errorf("cannot parse duration: %s", chain.ListenerBackOffDuration)
		}

		msg.Configuration.Chains[chain.UUID] = &dynamic.Chain{
			UUID:     chain.UUID,
			TenantID: chain.TenantID,
			OwnerID:  chain.OwnerID,
			Name:     chain.Name,
			URL:      utils.GetProxyURL(p.conf.ProxyURL, chain.UUID),
			ChainID:  chain.ChainID.String(),
			Listener: dynamic.Listener{
				StartingBlock:     chain.ListenerStartingBlock,
				CurrentBlock:      chain.ListenerCurrentBlock,
				Depth:             chain.ListenerDepth,
				Backoff:           duration,
				ExternalTxEnabled: chain.ListenerExternalTxEnabled,
			},
		}
	}

	return msg
}
