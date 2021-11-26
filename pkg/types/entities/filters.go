package entities

import (
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type JobFilters struct {
	TxHashes      []ethcommon.Hash `validate:"omitempty,unique,dive"`
	ChainUUID     string           `validate:"omitempty,uuid"`
	Status        JobStatus        `validate:"omitempty,isJobStatus"`
	UpdatedAfter  time.Time        `validate:"omitempty"`
	ParentJobUUID string           `validate:"omitempty"`
	OnlyParents   bool             `validate:"omitempty"`
	WithLogs      bool             `validate:"omitempty"`
}

type TransactionRequestFilters struct {
	IdempotencyKeys []string `validate:"omitempty,unique"`
}

type FaucetFilters struct {
	Names     []string `validate:"omitempty,unique"`
	ChainRule string   `validate:"omitempty"`
	TenantID  string   `validate:"omitempty"`
}

type AccountFilters struct {
	Aliases  []string `validate:"omitempty,unique"`
	TenantID string   `validate:"omitempty"`
}

type ChainFilters struct {
	Names    []string `validate:"omitempty,unique"`
	TenantID string   `validate:"omitempty"`
}
