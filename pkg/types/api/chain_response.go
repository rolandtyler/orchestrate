package api

import (
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
)

type ChainResponse struct {
	UUID                      string                     `json:"uuid" example:"b4374e6f-b28a-4bad-b4fe-bda36eaf849c"`
	Name                      string                     `json:"name" example:"mainnet"`
	TenantID                  string                     `json:"tenantID" example:"tenant"`
	URLs                      []string                   `json:"urls" example:"https://mainnet.infura.io/v3/a73136601e6f4924a0baa4ed880b535e"`
	ChainID                   string                     `json:"chainID" example:"1"`
	ListenerDepth             uint64                     `json:"listenerDepth" example:"0"`
	ListenerCurrentBlock      uint64                     `json:"listenerCurrentBlock" example:"0"`
	ListenerStartingBlock     uint64                     `json:"listenerStartingBlock" example:"5000"`
	ListenerBackOffDuration   string                     `json:"listenerBackOffDuration" example:"5s"`
	ListenerExternalTxEnabled bool                       `json:"listenerExternalTxEnabled" example:"false"`
	PrivateTxManager          *entities.PrivateTxManager `json:"privateTxManager,omitempty"`
	CreatedAt                 time.Time                  `json:"createdAt" example:"2020-07-09T12:35:42.115395Z"`
	UpdatedAt                 time.Time                  `json:"updatedAt" example:"2020-07-09T12:35:42.115395Z"`
}
