package testutils

import (
	"math/big"
	"time"

	"github.com/consensys/orchestrate/pkg/toolkit/ethclient/rpc"
	"github.com/consensys/orchestrate/pkg/types/entities"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/consensys/orchestrate/pkg/utils"
)

func ParseIArray(args ...interface{}) (ret []interface{}) {
	ret = make([]interface{}, len(args))
	copy(ret, args)
	return
}

func FakeETHTransaction() *entities.ETHTransaction {
	return &entities.ETHTransaction{
		From:        FakeAddress(),
		To:          FakeAddress(),
		Nonce:       utils.ToPtr(hexutil.Uint64(1)).(*hexutil.Uint64),
		Value:       utils.ToPtr(hexutil.Big(*big.NewInt(50000))).(*hexutil.Big),
		GasPrice:    utils.ToPtr(hexutil.Big(*big.NewInt(10000))).(*hexutil.Big),
		Gas:         utils.ToPtr(hexutil.Uint64(21000)).(*hexutil.Uint64),
		Data:        hexutil.MustDecode("0x0"),
		Raw:         hexutil.MustDecode("0x0"),
		PrivateFrom: []byte("A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo="),
		PrivateFor:  [][]byte{[]byte("A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo="), []byte("B1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo=")},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func FakeETHTransactionParams() *entities.ETHTransactionParams {
	return &entities.ETHTransactionParams{
		From:            utils.ToPtr(ethcommon.HexToAddress("0x7357589f8e367c2C31F51242fB77B350A11830F3")).(*ethcommon.Address),
		To:              utils.ToPtr(ethcommon.HexToAddress("0x7357589f8e367c2C31F51242fB77B350A11830F2")).(*ethcommon.Address),
		Value:           utils.ToPtr(hexutil.Big(*big.NewInt(50000))).(*hexutil.Big),
		GasPrice:        utils.ToPtr(hexutil.Big(*big.NewInt(10000))).(*hexutil.Big),
		Gas:             utils.ToPtr(hexutil.Uint64(21000)).(*hexutil.Uint64),
		MethodSignature: "method(string,string)",
		Args:            ParseIArray("val1", "val2"),
		ContractName:    "ContractName",
		ContractTag:     "ContractTag",
		Nonce:           hexutil.Uint64(1),
	}
}

func FakeETHAccount() *entities.ETHAccount {
	return &entities.ETHAccount{
		Namespace:           "_",
		Address:             ethcommon.HexToAddress(utils.RandHexString(12)),
		PublicKey:           hexutil.MustDecode(utils.RandHexString(30)),
		CompressedPublicKey: hexutil.MustDecode(utils.RandHexString(20)),
	}
}

func FakeTesseraTransactionParams() *entities.ETHTransactionParams {
	tx := FakeETHTransactionParams()
	tx.PrivateFrom = []byte("ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc=")
	tx.PrivateFor = [][]byte{[]byte("ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bd=")}
	tx.Protocol = entities.TesseraChainType

	return tx
}

func FakeEEATransactionParams() *entities.ETHTransactionParams {
	tx := FakeETHTransactionParams()
	tx.PrivateFrom = []byte("ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Be=")
	tx.PrivacyGroupID = []byte("ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bf=")
	tx.Protocol = entities.EEAChainType

	return tx
}

func FakeRawTransactionParams() *entities.ETHTransactionParams {
	return &entities.ETHTransactionParams{
		Raw: hexutil.MustDecode("0xABCDE012312312"),
	}
}

func FakeTransferTransactionParams() *entities.ETHTransactionParams {
	from := ethcommon.HexToAddress("0x7357589f8e367c2C31F51242fB77B350A11830FA")
	return &entities.ETHTransactionParams{
		From:  &from,
		To:    utils.ToPtr(ethcommon.HexToAddress("0x7357589f8e367c2C31F51242fB77B350A11830FB")).(*ethcommon.Address),
		Value: utils.ToPtr(hexutil.Big(*big.NewInt(50000))).(*hexutil.Big),
	}
}

func FakeAddress() *ethcommon.Address {
	addr := ethcommon.HexToAddress(utils.RandHexString(20))
	return &addr
}

func FakeHash() hexutil.Bytes {
	return hexutil.MustDecode("0x" + utils.RandHexString(40))
}

func FakeFeeHistory(nextBaseFee *big.Int) *rpc.FeeHistory {
	result := &rpc.FeeHistory{}
	nBaseFee2 := hexutil.Big(*nextBaseFee)
	result.BaseFeePerGas = []hexutil.Big{nBaseFee2}
	return result
}
