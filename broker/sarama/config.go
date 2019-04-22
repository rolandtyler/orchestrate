package sarama

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault(kafkaAddressViperKey, kafkaAddressDefault)
	_ = viper.BindEnv(kafkaAddressViperKey, kafkaAddressEnv)
	viper.SetDefault(txCrafterViperKey, txCrafterTopicDefault)
	_ = viper.BindEnv(txCrafterViperKey, txCrafterTopicEnv)
	viper.SetDefault(txNonceViperKey, txNonceTopicDefault)
	_ = viper.BindEnv(txNonceViperKey, txNonceTopicEnv)
	viper.SetDefault(generateWalletViperKey, generateWalletDefault)
	_ = viper.BindEnv(generateWalletViperKey, generateWalletTopicEnv)
	viper.SetDefault(txSignerViperKey, txSignerTopicDefault)
	_ = viper.BindEnv(txSignerViperKey, txSignerTopicEnv)
	viper.SetDefault(txSenderViperKey, txSenderTopicDefault)
	_ = viper.BindEnv(txSenderViperKey, txSenderTopicEnv)
	viper.SetDefault(txDecoderViperKey, txDecoderTopicDefault)
	_ = viper.BindEnv(txDecoderViperKey, txDecoderTopicEnv)
	viper.SetDefault(txDecodedViperKey, txDecodedTopicDefault)
	_ = viper.BindEnv(txDecodedViperKey, txDecodedTopicEnv)
	viper.SetDefault(txRecoverViperKey, txRecoverTopicDefault)
	_ = viper.BindEnv(txRecoverViperKey, txRecoverTopicEnv)
	viper.SetDefault(kafkaGroupViperKey, kafkaGroupDefault)
	_ = viper.BindEnv(kafkaGroupViperKey, kafkaGroupEnv)
}

var (
	kafkaAddressFlag     = "kafka-address"
	kafkaAddressViperKey = "kafka.addresses"
	kafkaAddressDefault  = []string{"localhost:9092"}
	kafkaAddressEnv      = "KAFKA_ADDRESS"
)

// KafkaAddresses register flag for Kafka server addresses
func KafkaAddresses(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Address of Kafka server to connect to.
Environment variable: %q`, kafkaAddressEnv)
	f.StringSlice(kafkaAddressFlag, kafkaAddressDefault, desc)
	_ = viper.BindPFlag(kafkaAddressViperKey, f.Lookup(kafkaAddressFlag))
}

var (
	kafkaGroupFlag     = "kafka-group"
	kafkaGroupViperKey = "kafka.group"
	kafkaGroupDefault  = "group-e2e"
	kafkaGroupEnv      = "KAFKA_GROUP"
)

// KafkaGroup register flag for Kafka server addresses
func KafkaGroup(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Address of Kafka server to connect to.
Environment variable: %q`, kafkaGroupEnv)
	f.String(kafkaGroupFlag, kafkaGroupDefault, desc)
	_ = viper.BindPFlag(kafkaGroupViperKey, f.Lookup(kafkaGroupEnv))
}

var (
	txCrafterFlag         = "topic-crafter"
	txCrafterViperKey     = "kafka.topic.crafter"
	txCrafterTopicEnv     = "KAFKA_TOPIC_TX_CRAFTER"
	txCrafterTopicDefault = "topic-tx-crafter"

	txNonceFlag         = "topic-nonce"
	txNonceViperKey     = "kafka.topic.nonce"
	txNonceTopicEnv     = "KAFKA_TOPIC_TX_NONCE"
	txNonceTopicDefault = "topic-tx-nonce"

	generateWalletFlag     = "topic-wallet"
	generateWalletViperKey = "kafka.topic.wallet.generator"
	generateWalletTopicEnv = "KAFKA_TOPIC_WALLET_GENERATOR"
	generateWalletDefault  = "topic-wallet-generator"

	txSignerFlag         = "topic-signer"
	txSignerViperKey     = "kafka.topic.signer"
	txSignerTopicEnv     = "KAFKA_TOPIC_TX_SIGNER"
	txSignerTopicDefault = "topic-tx-signer"

	txSenderFlag         = "topic-sender"
	txSenderViperKey     = "kafka.topic.sender"
	txSenderTopicEnv     = "KAFKA_TOPIC_TX_SENDER"
	txSenderTopicDefault = "topic-tx-sender"

	txDecoderFlag         = "topic-decoder"
	txDecoderViperKey     = "kafka.topic.decoder"
	txDecoderTopicEnv     = "KAFKA_TOPIC_TX_DECODER"
	txDecoderTopicDefault = "topic-tx-decoder"

	txDecodedFlag         = "topic-decoded"
	txDecodedViperKey     = "kafka.topic.decoded"
	txDecodedTopicEnv     = "KAFKA_TOPIC_TX_DECODED"
	txDecodedTopicDefault = "topic-tx-decoded"

	txRecoverFlag         = "topic-recover"
	txRecoverViperKey     = "kafka.topic.recover"
	txRecoverTopicEnv     = "KAFKA_TOPIC_TX_RECOVER"
	txRecoverTopicDefault = "topic-tx-recover"
)

// TODO: implement test for all Topics flags & Goup flags

// KafkaTopicTxCrafter register flag for Kafka topic
func KafkaTopicTxCrafter(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to have transaction payload crafted
Environment variable: %q`, txCrafterTopicEnv)
	f.String(txCrafterFlag, txCrafterTopicDefault, desc)
	_ = viper.BindPFlag(txCrafterViperKey, f.Lookup(txCrafterFlag))
}

// KafkaTopicTxNonce register flag for Kafka topic
func KafkaTopicTxNonce(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to have transaction nonce set
Environment variable: %q`, txNonceViperKey)
	f.String(txNonceFlag, txNonceTopicDefault, desc)
	_ = viper.BindPFlag(txNonceViperKey, f.Lookup(txNonceFlag))
}

// KafkaTopicTxSigner register flag for Kafka topic
func KafkaTopicTxSigner(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to have transaction signed
Environment variable: %q`, txSignerViperKey)
	f.String(txSignerFlag, txSignerTopicDefault, desc)
	_ = viper.BindPFlag(txSignerViperKey, f.Lookup(txSignerFlag))
}

// KafkaTopicWalletGenerator register flag for Kafka topic
func KafkaTopicWalletGenerator(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to generate a new wallet
Environment variable: %q`, generateWalletViperKey)
	f.String(generateWalletFlag, generateWalletDefault, desc)
	_ = viper.BindPFlag(generateWalletViperKey, f.Lookup(generateWalletFlag))
}

// KafkaTopicTxSender register flag for Kafka topic
func KafkaTopicTxSender(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to have transaction sent
Environment variable: %q`, txSenderViperKey)
	f.String(txSenderFlag, txSenderTopicDefault, desc)
	_ = viper.BindPFlag(txSenderViperKey, f.Lookup(txSenderFlag))
}

// KafkaTopicTxDecoder register flag for Kafka topic
func KafkaTopicTxDecoder(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to have receipt decoded
Environment variable: %q`, txDecoderViperKey)
	f.String(txDecoderFlag, txDecoderTopicDefault, desc)
	_ = viper.BindPFlag(txDecoderViperKey, f.Lookup(txDecoderFlag))
}

// KafkaTopicTxRecover register flag for Kafka topic
func KafkaTopicTxRecover(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages waiting to have transaction recovered
Environment variable: %q`, txRecoverViperKey)
	f.String(txRecoverFlag, txRecoverTopicDefault, desc)
	_ = viper.BindPFlag(txRecoverViperKey, f.Lookup(txRecoverFlag))
}

// KafkaTopicTxDecoded register flag for Kafka topic
func KafkaTopicTxDecoded(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Kafka topic for messages which receipt has been decoded
Environment variable: %q`, txDecodedViperKey)
	f.String(txDecodedFlag, txDecodedTopicDefault, desc)
	_ = viper.BindPFlag(txDecodedViperKey, f.Lookup(txDecodedFlag))
}

// Kafka Consumer group environment variables
var (
	crafterGroupFlag     = "group-crafter"
	crafterGroupViperKey = "kafka.group.crafter"
	crafterGroupEnv      = "KAFKA_GROUP_CRAFTER"
	crafterGroupDefault  = "group-crafter"

	nonceGroupFlag     = "group-nonce"
	nonceGroupViperKey = "kafka.group.nonce"
	nonceGroupEnv      = "KAFKA_GROUP_NONCE"
	nonceGroupDefault  = "group-nonce"

	signerGroupFlag     = "group-signer"
	signerGroupViperKey = "kafka.group.signer"
	signerGroupEnv      = "KAFKA_GROUP_SIGNER"
	signerGroupDefault  = "group-signer"

	senderGroupFlag     = "group-sender"
	senderGroupViperKey = "kafka.group.sender"
	senderGroupEnv      = "KAFKA_GROUP_SENDER"
	senderGroupDefault  = "group-sender"

	decoderGroupFlag     = "group-decoder"
	decoderGroupViperKey = "kafka.group.decoder"
	decoderGroupEnv      = "KAFKA_GROUP_DECODER"
	decoderGroupDefault  = "group-decoder"

	bridgeGroupFlag     = "group-bridge"
	bridgeGroupViperKey = "kafka.group.bridge"
	bridgeGroupEnv      = "KAFKA_GROUP_BRIDGE"
	bridgeGroupDefault  = "tx-group-bridge"
)

// consumerGroup register flag for a kafka consumer group
func consumerGroup(f *pflag.FlagSet, flag, key, env, defaultValue string) {
	desc := fmt.Sprintf(`Kafka consumer group name
Environment variable: %q`, env)
	f.String(flag, defaultValue, desc)
	_ = viper.BindPFlag(key, f.Lookup(flag))
}

// CrafterGroup register flag for kafka crafter group
func CrafterGroup(f *pflag.FlagSet) {
	consumerGroup(f, crafterGroupFlag, crafterGroupViperKey, crafterGroupEnv, crafterGroupDefault)
}

// NonceGroup register flag for kafka nonce group
func NonceGroup(f *pflag.FlagSet) {
	consumerGroup(f, nonceGroupFlag, nonceGroupViperKey, nonceGroupEnv, nonceGroupDefault)
}

// SignerGroup register flag for kafka signer group
func SignerGroup(f *pflag.FlagSet) {
	consumerGroup(f, signerGroupFlag, signerGroupViperKey, signerGroupEnv, signerGroupDefault)
}

// SenderGroup register flag for kafka sender group
func SenderGroup(f *pflag.FlagSet) {
	consumerGroup(f, senderGroupFlag, senderGroupViperKey, senderGroupEnv, senderGroupDefault)
}

// DecoderGroup register flag for kafka decoder group
func DecoderGroup(f *pflag.FlagSet) {
	consumerGroup(f, decoderGroupFlag, decoderGroupViperKey, decoderGroupEnv, decoderGroupDefault)
}

// BridgeGroup register flag for kafka decoder group
func BridgeGroup(f *pflag.FlagSet) {
	consumerGroup(f, bridgeGroupFlag, bridgeGroupViperKey, bridgeGroupEnv, bridgeGroupDefault)
}
