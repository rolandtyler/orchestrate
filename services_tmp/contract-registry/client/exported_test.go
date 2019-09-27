package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	svc "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/services/contract-registry"
)

func TestInit(t *testing.T) {
	Init(context.Background())
	assert.NotNil(t, GlobalContractRegistryClient(), "Global should have been set")

	var c svc.RegistryClient
	SetGlobalContractRegistryClient(c)
	assert.Nil(t, GlobalContractRegistryClient(), "Global should be reset to nil")
}
