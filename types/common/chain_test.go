package common

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	chain := &Chain{
		Id:       "42",
		IsEIP155: true,
	}
	assert.Equal(t, int64(42), chain.ID().Int64(), "Chain ID should match")

	chain.SetID(big.NewInt(54))
	assert.Equal(t, "54", chain.Id, "Chain ID should have been properly set match")
}
