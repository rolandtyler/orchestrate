package entities

import (
	"time"

	"github.com/consensys/quorum/common/hexutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Faucet struct {
	UUID            string
	Name            string
	TenantID        string
	ChainRule       string
	CreditorAccount ethcommon.Address
	MaxBalance      hexutil.Big
	Amount          hexutil.Big
	Cooldown        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
