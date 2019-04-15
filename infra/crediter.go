package infra

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	infSarama "gitlab.com/ConsenSys/client/fr/core-stack/infra/sarama.git"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/core/services"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/ethereum"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/trace"
)

// SaramaCrediter allows to credit by sending messages to a Kafka topic
type SaramaCrediter struct {
	addresses map[string]ethcommon.Address

	p sarama.SyncProducer
	m *infSarama.Marshaller
}

func parseAddresses(addresses []string) (map[string]ethcommon.Address, error) {
	m := make(map[string]ethcommon.Address)
	for _, addr := range addresses {
		split := strings.Split(addr, "@")
		if len(split) != 2 {
			return nil, fmt.Errorf("Could not parse faucet address %q (expected format %q)", addr, "<address>@<chainID>")
		}

		if !ethcommon.IsHexAddress(split[0]) {
			return nil, fmt.Errorf("Invalid Ethereum address %q", split[0])
		}

		chain, ok := new(big.Int).SetString(split[1], 10)
		if !ok {
			log.Fatal("Could not cast chain ID to big.Int")
		}
		m[hexutil.EncodeBig(chain)] = ethcommon.HexToAddress(split[0])
	}
	return m, nil
}

// NewSaramaCrediter creates a new SaramaCrediter
func NewSaramaCrediter(p sarama.SyncProducer) (*SaramaCrediter, error) {
	addresses, err := parseAddresses(viper.GetStringSlice("faucet.addresses"))
	if err != nil {
		return nil, err
	}

	return &SaramaCrediter{
		addresses: addresses,
		p:         p,
		m:         infSarama.NewMarshaller(),
	}, nil
}

// Credit credit a given request by sending a message to a Kafka topic
func (c *SaramaCrediter) Credit(ctx context.Context, r *services.FaucetRequest) (*big.Int, bool, error) {
	// Prepare Faucet message
	msg, err := c.PrepareFaucetMsg(r)
	if err != nil {
		return big.NewInt(0), false, err
	}

	// Send message
	_, _, err = c.p.SendMessage(&msg)
	if err != nil {
		return big.NewInt(0), false, err
	}

	return r.Value, true, nil
}

// PrepareFaucetMsg creates a credit message to send to a specific topic
func (c *SaramaCrediter) PrepareFaucetMsg(r *services.FaucetRequest) (sarama.ProducerMessage, error) {
	// Determine Address of the faucet for requested chain
	faucetAddress := c.addresses[hexutil.EncodeBig(r.ChainID)]

	if (faucetAddress != ethcommon.Address{}) {
		// Create Trace for Crediting message
		faucetTrace := &trace.Trace{
			Metadata: &trace.Metadata{Id: uuid.NewV4().String()},
			Chain:    &common.Chain{Id: hexutil.EncodeBig(r.ChainID)},
			Sender:   &common.Account{Addr: faucetAddress.String()},
			Tx: &ethereum.Transaction{TxData: &ethereum.TxData{
				To:    r.Address.String(),
				Value: hexutil.EncodeBig(r.Value),
			}},
		}

		// Create Producer message
		var msg sarama.ProducerMessage
		err := c.m.Marshal(faucetTrace, &msg)
		if err != nil {
			return sarama.ProducerMessage{}, err
		}
		msg.Topic = viper.GetString("faucet.topic")

		return msg, nil
	}

	return sarama.ProducerMessage{}, fmt.Errorf("crediter: No faucet address available for ChainId %v", hexutil.EncodeBig(r.ChainID))
}
