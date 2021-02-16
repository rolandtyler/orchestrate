// +build unit

package session

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/account"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"
)

// Test against hash and rlpEncoded signature from
func TestExecuteForTx(t *testing.T) {
	// crtl := gomock.NewController(t)
	multitenancy.Init(context.Background())
	acc := account.NewAccount()
	_ = acc.FromPrivateKey("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63")

	s := NewSigningSession()
	_ = s.SetChain(big.NewInt(2018))
	_ = s.SetAccount(acc)

	h, _ := hexutil.Decode("0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610221806100606000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633fa4f2451461005c5780636057361d1461008757806367e404ce146100b4575b600080fd5b34801561006857600080fd5b5061007161010b565b6040518082815260200191505060405180910390f35b34801561009357600080fd5b506100b260048036038101908080359060200190929190505050610115565b005b3480156100c057600080fd5b506100c96101cb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000600254905090565b7fc9db20adedc6cf2b5d25252b101ab03e124902a73fcb12b753f3d1aaa2d8f9f53382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a18060028190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050905600a165627a7a723058208efaf938851fb2d235f8bf9a9685f149129a30fe0f4b20a6c1885dc02f639eba002900000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000d5065676153797320546f6b656e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000075065676153797300000000000000000000000000000000000000000000000000")
	tx := ethtypes.NewContractCreation(
		0,
		big.NewInt(0),
		3000000,
		big.NewInt(0),
		h,
	)

	raw, hash, err := s.ExecuteForTx(tx)
	expectedSignedRlpEncoded, _ := hexutil.Decode("0xf903d18080832dc6c08080b90381608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610221806100606000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633fa4f2451461005c5780636057361d1461008757806367e404ce146100b4575b600080fd5b34801561006857600080fd5b5061007161010b565b6040518082815260200191505060405180910390f35b34801561009357600080fd5b506100b260048036038101908080359060200190929190505050610115565b005b3480156100c057600080fd5b506100c96101cb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000600254905090565b7fc9db20adedc6cf2b5d25252b101ab03e124902a73fcb12b753f3d1aaa2d8f9f53382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a18060028190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050905600a165627a7a723058208efaf938851fb2d235f8bf9a9685f149129a30fe0f4b20a6c1885dc02f639eba002900000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000d5065676153797320546f6b656e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000075065676153797300000000000000000000000000000000000000000000000000820fe8a046603f798e67638c724275508aa8cb0074a0adc3e8c97279858d4b0d7404d19fa06fed81c1cec0a76c7ef0e8a5a67b60d6a3df9ac2edb21eb5bb9da832248d4fb5")
	expectedHash := "0xf8153ed0bf7747b698f18595337b1d08efd9ffaa79132585ae1be0211244630a"
	assert.Equal(t, expectedSignedRlpEncoded, raw)
	assert.Equal(t, expectedHash, hash.Hex())
	assert.NoError(t, err)
}

// Test against hash and rlpEncoded signature from
// https://github.com/PegaSysEng/web3js-eea/blob/master/example/erc20.js
func TestExecuteForEEATx(t *testing.T) {
	// crtl := gomock.NewController(t)
	multitenancy.Init(context.Background())
	acc := account.NewAccount()
	_ = acc.FromPrivateKey("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63")

	s := NewSigningSession()
	_ = s.SetChain(big.NewInt(2018))
	_ = s.SetAccount(acc)

	h, _ := hexutil.Decode("0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610221806100606000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633fa4f2451461005c5780636057361d1461008757806367e404ce146100b4575b600080fd5b34801561006857600080fd5b5061007161010b565b6040518082815260200191505060405180910390f35b34801561009357600080fd5b506100b260048036038101908080359060200190929190505050610115565b005b3480156100c057600080fd5b506100c96101cb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000600254905090565b7fc9db20adedc6cf2b5d25252b101ab03e124902a73fcb12b753f3d1aaa2d8f9f53382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a18060028190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050905600a165627a7a723058208efaf938851fb2d235f8bf9a9685f149129a30fe0f4b20a6c1885dc02f639eba002900000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000d5065676153797320546f6b656e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000075065676153797300000000000000000000000000000000000000000000000000")
	tx := ethtypes.NewContractCreation(
		0,
		big.NewInt(0),
		3000000,
		big.NewInt(0),
		h,
	)
	priv := &types.PrivateArgs{
		PrivateFrom:   "A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo=",
		PrivateFor:    []string{"Ko2bVqD+nNlNYL5EE7y3IdOnviftjiizpjRt+HTuFBs="},
		PrivateTxType: "restricted",
	}

	raw, hash, err := s.ExecuteForEEATx(tx, priv)

	expectedSignedRlpEncoded, _ := hexutil.Decode("0xf9041f8080832dc6c08080b90381608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610221806100606000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633fa4f2451461005c5780636057361d1461008757806367e404ce146100b4575b600080fd5b34801561006857600080fd5b5061007161010b565b6040518082815260200191505060405180910390f35b34801561009357600080fd5b506100b260048036038101908080359060200190929190505050610115565b005b3480156100c057600080fd5b506100c96101cb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000600254905090565b7fc9db20adedc6cf2b5d25252b101ab03e124902a73fcb12b753f3d1aaa2d8f9f53382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a18060028190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050905600a165627a7a723058208efaf938851fb2d235f8bf9a9685f149129a30fe0f4b20a6c1885dc02f639eba002900000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000d5065676153797320546f6b656e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000075065676153797300000000000000000000000000000000000000000000000000820fe8a0ce0855abf94c357db638cf606fdbe235f121df38205ba9fe298c1b21d9d52d78a0310afc9e284428eda8820863230d7d0bc0dacb1010a3ce7679d60402375852bfa0035695b4cc4b0941e60551d7a19cf30603db5bfc23e5ac43a56f57f25f75486ae1a02a8d9b56a0fe9cd94d60be4413bcb721d3a7be27ed8e28b3a6346df874ee141b8a72657374726963746564")
	expectedHash := "0x03ea3d334980b5761c836e6b0bc3e84321316520a024cfaa454a7b8e88c7ccc1"
	assert.Equal(t, expectedSignedRlpEncoded, raw)
	assert.Equal(t, expectedHash, hash.Hex())
	assert.NoError(t, err)
}