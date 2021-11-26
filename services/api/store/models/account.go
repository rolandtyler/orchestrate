package models

import (
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Account struct {
	tableName struct{} `pg:"accounts"` // nolint:unused,structcheck // reason

	ID                  int
	Alias               string
	Address             string
	PublicKey           string
	CompressedPublicKey string
	TenantID            string
	OwnerID             string
	Attributes          map[string]string
	// TODO add internal labels to store accountID

	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time `pg:"default:now()"`
}
