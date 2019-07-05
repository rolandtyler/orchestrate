package generator

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	encoding "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/encoding/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/handlers/producer"
)

// PrepareMsg prepare message to produce from TxContexts
func PrepareMsg(txctx *engine.TxContext, msg *sarama.ProducerMessage) error {
	// Marshal Envelope into sarama Message
	err := encoding.Marshal(txctx.Envelope, msg)
	if err != nil {
		return err
	}

	msg.Topic = viper.GetString("kafka.topic.wallet.generated")
	return nil
}

// Producer creates a producer handler
func Producer(p sarama.SyncProducer) engine.HandlerFunc {
	return producer.Producer(p, PrepareMsg)
}
