package signer

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/ConsenSys/client/fr/core-stack/service/ethereum.git/tessera"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/ethereum.git/types"

	"github.com/magiconair/properties/assert"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/chain"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/ethereum"
	eeaHandlers "gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-signer.git/handlers/vault/signer/eea"
	ethereumHandlers "gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-signer.git/handlers/vault/signer/ethereum"
	tesseraHandlers "gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-signer.git/handlers/vault/signer/tessera"
)

type MockTxSigner struct {
	t *testing.T
}

type MockTesseraClient struct {
	t *testing.T
}

var alreadySignedTx = "0x00"
var signedTx = "0x01"
var signedPrivateTx = "0x02"
var signedTesseraTx = "0x03"

func (s *MockTxSigner) SignTx(netChain *chain.Chain, a ethcommon.Address, tx *ethtypes.Transaction) (raw []byte, hash *ethcommon.Hash, err error) {
	if netChain.ID().String() == "0" {
		return []byte(``), nil, fmt.Errorf("could not sign public ethereum transaction")
	}
	h := ethcommon.HexToHash("0xabcdef")
	return hexutil.MustDecode(signedTx), &h, nil
}

func (s *MockTxSigner) SignPrivateEEATx(netChain *chain.Chain, a ethcommon.Address, tx *ethtypes.Transaction, privateArgs *types.PrivateArgs) (raw []byte, hash *ethcommon.Hash, err error) {
	if netChain.ID().String() == "0" {
		return []byte(``), nil, fmt.Errorf("could not sign eea transaction")
	}
	h := ethcommon.HexToHash("0xabcdef")
	return hexutil.MustDecode(signedPrivateTx), &h, nil
}

func (s *MockTxSigner) SignPrivateTesseraTx(netChain *chain.Chain, a ethcommon.Address, tx *ethtypes.Transaction) (raw []byte, txHash *ethcommon.Hash, err error) {
	if netChain.ID().String() == "0" {
		return []byte(``), nil, fmt.Errorf("could not sign tessera transaction")
	}
	h := ethcommon.HexToHash("0xabcdef")
	return hexutil.MustDecode(signedTesseraTx), &h, nil
}

func (s *MockTxSigner) SignMsg(a ethcommon.Address, msg string) (rsv []byte, hash *ethcommon.Hash, err error) {
	return []byte{}, nil, fmt.Errorf("signMsg not implemented")
}

func (s *MockTxSigner) GenerateWallet() (add *ethcommon.Address, err error) {
	return nil, fmt.Errorf("signMsg not implemented")
}

func (s *MockTxSigner) SignRawHash(a ethcommon.Address, hash []byte) (rsv []byte, err error) {
	return []byte{}, fmt.Errorf("signMsg not implemented")
}

func (s *MockTxSigner) ImportPrivateKey(priv string) (err error) {
	return fmt.Errorf("importPrivateKey not implemented")
}

func (tc *MockTesseraClient) AddClient(chainID string, tesseraEndpoint tessera.EnclaveEndpoint) {

}

func (tc *MockTesseraClient) StoreRaw(chainID string, rawTx []byte, privateFrom string) (txHash []byte, err error) {
	if chainID == "0" {
		return []byte(``), fmt.Errorf("mock: store raw failed")
	}
	return hexutil.MustDecode("0xabcdef"), nil
}

func (tc *MockTesseraClient) GetStatus(chainID string) (status string, err error) {
	if chainID == "0" {
		return "", fmt.Errorf("mock: get status failed")
	}
	return "", nil
}

func makeSignerContext(i int) *engine.TxContext {
	txctx := engine.NewTxContext()
	txctx.Reset()
	txctx.Logger = log.NewEntry(log.StandardLogger())

	switch i % 8 {
	case 0:
		h := ethcommon.HexToHash("0x12345678")
		txctx.Envelope.Chain = chain.CreateChainInt(10)
		txctx.Envelope.Tx = &ethereum.Transaction{
			Raw:  ethereum.HexToData(alreadySignedTx),
			Hash: ethereum.CreateHash(h.Bytes()),
		}
		txctx.Set("errors", 0)
		txctx.Set("raw", alreadySignedTx)
		txctx.Set("hash", "0x0000000000000000000000000000000000000000000000000000000012345678")
	case 1:
		h := ethcommon.HexToHash("0x12345678")
		txctx.Envelope.Chain = chain.CreateChainInt(0)
		txctx.Envelope.Tx = &ethereum.Transaction{
			Raw:  ethereum.HexToData(alreadySignedTx),
			Hash: ethereum.CreateHash(h.Bytes()),
		}

		txctx.Set("errors", 0)
		txctx.Set("raw", alreadySignedTx)
		txctx.Set("hash", "0x0000000000000000000000000000000000000000000000000000000012345678")
	case 2:
		txctx.Envelope.Chain = chain.CreateChainInt(0)
		txctx.Envelope.Tx = &ethereum.Transaction{}
		txctx.Set("errors", 1)
		txctx.Set("raw", "0x")
		txctx.Set("hash", "0x")
	case 3:
		txctx.Envelope.Chain = chain.CreateChainInt(10)
		txctx.Envelope.Tx = &ethereum.Transaction{}
		txctx.Set("errors", 0)
		txctx.Set("raw", signedTx)
		txctx.Set("hash", "0x0000000000000000000000000000000000000000000000000000000000abcdef")
	case 4:
		txctx.Envelope.Chain = chain.CreateChainInt(10)
		txctx.Envelope.Tx = &ethereum.Transaction{
			TxData: &ethereum.TxData{
				Data: &ethereum.Data{
					Raw: []byte{0},
				},
			},
		}
		txctx.Envelope.Protocol = &chain.Protocol{
			Type: chain.ProtocolType_QUORUM_TESSERA,
		}
		txctx.Set("errors", 0)
		txctx.Set("raw", signedTesseraTx)
		txctx.Set("hash", "0x0000000000000000000000000000000000000000000000000000000000abcdef")
	case 5:
		txctx.Envelope.Chain = chain.CreateChainInt(10)
		txctx.Envelope.Tx = &ethereum.Transaction{}
		txctx.Envelope.Protocol = &chain.Protocol{
			Type: chain.ProtocolType_QUORUM_TESSERA,
		}
		txctx.Set("errors", 1)
		txctx.Set("raw", "0x")
		txctx.Set("hash", "0x")
	case 6:
		txctx.Envelope.Chain = chain.CreateChainInt(0)
		txctx.Envelope.Tx = &ethereum.Transaction{
			TxData: &ethereum.TxData{
				Data: &ethereum.Data{
					Raw: []byte{0},
				},
			},
		}
		txctx.Envelope.Protocol = &chain.Protocol{
			Type: chain.ProtocolType_QUORUM_TESSERA,
		}
		txctx.Set("errors", 1)
		txctx.Set("raw", "0x")
		txctx.Set("hash", "0x")
	case 7:
		txctx.Envelope.Chain = chain.CreateChainInt(10)
		txctx.Envelope.Tx = &ethereum.Transaction{}
		txctx.Envelope.Protocol = &chain.Protocol{
			Type: chain.ProtocolType_PANTHEON_ORION,
		}
		txctx.Set("errors", 0)
		txctx.Set("raw", signedPrivateTx)
		txctx.Set("hash", "0x0000000000000000000000000000000000000000000000000000000000abcdef")
	}
	return txctx
}

func TestSigner(t *testing.T) {

	s := &MockTxSigner{t: t}
	tc := &MockTesseraClient{t: t}

	signer := TxSigner(
		eeaHandlers.Signer(s),
		ethereumHandlers.Signer(s),
		tesseraHandlers.Signer(s, tc),
	)

	rounds := 25
	outs := make(chan *engine.TxContext, rounds)
	wg := &sync.WaitGroup{}
	for i := 0; i < rounds; i++ {
		wg.Add(1)
		txctx := makeSignerContext(i)
		go func(txctx *engine.TxContext) {
			defer wg.Done()
			signer(txctx)
			outs <- txctx
		}(txctx)
	}
	wg.Wait()
	close(outs)

	if len(outs) != rounds {
		t.Errorf("Signer: expected %v outs but got %v", rounds, len(outs))
	}

	for out := range outs {
		errCount, raw, hash := out.Get("errors").(int), out.Get("raw").(string), out.Get("hash").(string)
		assert.Equal(t, len(out.Envelope.Errors), errCount, fmt.Sprintf("Signer: expected %v errors but got %v, the TxContext was: ", errCount, out.Envelope.Errors))
		assert.Equal(t, out.Envelope.Tx.GetRaw().Hex(), raw, fmt.Sprintf("Signer: expected Raw %v but got %v", raw, out.Envelope.Tx.GetRaw().Hex()))
		assert.Equal(t, out.Envelope.Tx.GetHash().Hex(), hash, fmt.Sprintf("Signer: expected hash %v but got %v", hash, out.Envelope.Tx.GetHash().Hex()))
	}
}
