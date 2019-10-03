package txnonce

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/types/envelope"
)

func TestPrepareMsg(t *testing.T) {
	// No error
	txctx := engine.NewTxContext()
	txctx.Envelope = &envelope.Envelope{}
	msg := &sarama.ProducerMessage{}
	_ = PrepareMsg(txctx, msg)
	assert.Equal(t, "topic-tx-signer", msg.Topic, "If no error out topic should be nonce")

	// Faucet warning
	_ = txctx.Error(errors.Warningf("warning"))
	_ = PrepareMsg(txctx, msg)
	assert.Equal(t, "topic-tx-signer", msg.Topic, "If Faucet warning out topic should be nonce")

	// Classic error
	_ = txctx.Error(errors.ConnectionError("Connection error"))
	_ = PrepareMsg(txctx, msg)
	assert.Equal(t, "topic-tx-recover", msg.Topic, "If error out topic should be recovery")
}
