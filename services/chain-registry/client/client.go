package client

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/types"
)

type ChainRegistryClient interface {
	GetChains(ctx context.Context) ([]*types.Chain, error)
	GetChainByName(ctx context.Context, chainName string) (*types.Chain, error)
	GetChainByUUID(ctx context.Context, chainUUID string) (*types.Chain, error)
	UpdateBlockPosition(ctx context.Context, chainUUID string, blockNumber int64) error
}
