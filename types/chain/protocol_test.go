package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtocol(t *testing.T) {
	p := Protocol{Type: ProtocolType_PANTHEON_ORION}
	assert.Truef(t, p.IsPantheon(), "IsPantheon should be true")

	p = Protocol{Type: ProtocolType_QUORUM_CONSTELLATION}
	assert.Truef(t, p.IsConstellation(), "IsConstellation should be true")

	p = Protocol{Type: ProtocolType_QUORUM_TESSERA}
	assert.Truef(t, p.IsTessera(), "IsQuorum should be true")

	p = Protocol{Type: ProtocolType_ETHEREUM_CONSTANTINOPLE}
	assert.Falsef(t, p.IsTessera(), "IsTessera should be false")
	assert.Falsef(t, p.IsConstellation(), "IsConstellation should be false")
	assert.Falsef(t, p.IsPantheon(), "IsPantheon should be false")
}