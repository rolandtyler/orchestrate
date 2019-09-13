package mock

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/errors"
	svc "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/services/contract-registry"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/abi"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/chain"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/common"
	ierror "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/error"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/ethereum"
)

var ERC20 = []byte(
	`[{
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "account", "type": "address"},
      {"indexed": false, "name": "account2", "type": "address"}
    ],
    "name": "MinterAdded",
    "type": "event"
  },
  {
    "constant": true,
    "inputs": [
      {"name": "account", "type": "address"}
    ],
    "name": "isMinter",
    "outputs": [
      {"name": "", "type": "bool"}
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
    }]`)

var ERC20bis = []byte(
	`[{
	"anonymous": false,
	"inputs": [
	  {"indexed": false, "name": "account", "type": "address"},
	  {"indexed": true, "name": "account2", "type": "address"}
	],
	"name": "MinterAdded",
	"type": "event"
  },
  {
	"anonymous": false,
	"inputs": [
	  {"indexed": false, "name": "account", "type": "address"},
	  {"indexed": true, "name": "account2", "type": "address"}
	],
	"name": "MinterAddedBis",
	"type": "event"
  },
  {
	"constant": true,
	"inputs": [
	  {"name": "account", "type": "address"}
	],
	"name": "isMinter",
	"outputs": [
	  {"name": "", "type": "bool"}
	],
	"payable": false,
	"stateMutability": "view",
	"type": "function"
	}]`)

var methodSig = []byte("isMinter(address)")
var eventSig = []byte("MinterAdded(address,address)")

var ERC20Contract = &abi.Contract{
	Id: &abi.ContractId{
		Name: "ERC20",
		Tag:  "v1.0.0",
	},
	Abi:              ERC20,
	Bytecode:         []byte{1, 2},
	DeployedBytecode: []byte{1, 2, 3},
}
var ERC20ContractBis = &abi.Contract{
	Id: &abi.ContractId{
		Name: "ERC20",
		Tag:  "v1.0.1",
	},
	Abi:              ERC20bis,
	Bytecode:         []byte{1, 3},
	DeployedBytecode: []byte{1, 2, 4},
}

var methodJSONs, eventJSONs, _ = parseRawJSON(ERC20Contract.Abi)
var _, eventJSONsBis, _ = parseRawJSON(ERC20ContractBis.Abi)

var ContractInstance = common.AccountInstance{
	Chain:   &chain.Chain{Id: big.NewInt(3).Bytes()},
	Account: ethereum.HexToAccount("0xBA826fEc90CEFdf6706858E5FbaFcb27A290Fbe0"),
}

func TestRegisterContract(t *testing.T) {
	r := NewRegistry()
	_, err := r.RegisterContract(
		context.Background(),
		&svc.RegisterContractRequest{
			Contract: &abi.Contract{
				Id: &abi.ContractId{
					Name: "ERC20",
					Tag:  "v1.0.0",
				},
				Abi: []byte{},
			},
		},
	)
	assert.NoError(t, err, "Should not error on empty things")

	_, err = r.RegisterContract(context.Background(),
		&svc.RegisterContractRequest{Contract: ERC20Contract},
	)
	assert.NoError(t, err, "Should register contract properly")

	_, err = r.RegisterContract(context.Background(),
		&svc.RegisterContractRequest{Contract: ERC20Contract},
	)
	assert.NoError(t, err, "Should register contract properly twice")
}

func TestContractRegistryBySig(t *testing.T) {
	r := NewRegistry()
	_, err := r.RegisterContract(context.Background(),
		&svc.RegisterContractRequest{Contract: ERC20Contract},
	)
	assert.NoError(t, err)
	_, err = r.RegisterContract(context.Background(),
		&svc.RegisterContractRequest{Contract: ERC20ContractBis},
	)
	assert.NoError(t, err)

	// Get Contract
	contractResp, err := r.GetContract(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "v1.0.0",
			},
		})
	assert.NoError(t, err)
	assert.Equal(t, ERC20Contract.Abi, contractResp.GetContract().GetAbi())

	abiResp, err := r.GetContractABI(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "covfefe",
			},
		})
	assert.Error(t, err, "GetContractABI should error when unknown contract")
	ierr, ok := err.(*ierror.Error)
	assert.True(t, ok, "GetContractABI error should cast to internal error")
	assert.Equal(t, "contract-registry.mock", ierr.GetComponent(), "GetContractABI error component should be correct")
	assert.True(t, errors.IsStorageError(ierr), "GetContractABI error should be a storage error")
	assert.Nil(t, abiResp)

	// Get ABI
	abiResp, err = r.GetContractABI(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "v1.0.0",
			},
		})
	assert.NoError(t, err)
	assert.Equal(t, ERC20Contract.Abi, abiResp.GetAbi())

	abiResp, err = r.GetContractABI(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "covfefe",
			},
		})
	assert.Error(t, err, "GetContractABI should error when unknown contract")
	ierr, ok = err.(*ierror.Error)
	assert.True(t, ok, "GetContractABI error should cast to internal error")
	assert.Equal(t, "contract-registry.mock", ierr.GetComponent(), "GetContractABI error component should be correct")
	assert.True(t, errors.IsStorageError(ierr), "GetContractABI error should be a storage error")
	assert.Nil(t, abiResp)

	// Get Bytecode
	bytecodeResp, err := r.GetContractBytecode(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "v1.0.0",
			},
		})
	assert.NoError(t, err)
	assert.Equal(t, ERC20Contract.Bytecode, bytecodeResp.GetBytecode())
	bytecodeResp, err = r.GetContractBytecode(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "covfefe",
			},
		})
	assert.Error(t, err, "GetContractBytecode should error when unknown contract")
	ierr, ok = err.(*ierror.Error)
	assert.True(t, ok, "GetContractBytecode error should cast to internal error")
	assert.Equal(t, "contract-registry.mock", ierr.GetComponent(), "GetContractBytecode error component should be correct")
	assert.True(t, errors.IsStorageError(ierr), "GetContractBytecode error should be a storage error")
	assert.Nil(t, bytecodeResp)

	// Get DeployedBytecode
	deployedBytecodeResp, err := r.GetContractDeployedBytecode(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "v1.0.0",
			},
		})
	assert.NoError(t, err)
	assert.Equal(t, ERC20Contract.DeployedBytecode, deployedBytecodeResp.GetDeployedBytecode())
	deployedBytecodeResp, err = r.GetContractDeployedBytecode(context.Background(),
		&svc.GetContractRequest{
			ContractId: &abi.ContractId{
				Name: "ERC20",
				Tag:  "covfefe",
			},
		})
	assert.Error(t, err, "Should error when unknown contract")
	ierr, ok = err.(*ierror.Error)
	assert.True(t, ok, "GetContractDeployedBytecode should cast to internal error")
	assert.Equal(t, "contract-registry.mock", ierr.GetComponent(), "GetContractDeployedBytecode error component should be correct")
	assert.True(t, errors.IsStorageError(ierr), "GetContractDeployedBytecode error should be a storage error")
	assert.Nil(t, deployedBytecodeResp)

	// Get Catalog
	namesResp, err := r.GetCatalog(context.Background(), &svc.GetCatalogRequest{})
	assert.NoError(t, err)
	assert.Equal(t, []string{"ERC20"}, namesResp.GetNames())

	// Get Tags
	tagsResp, err := r.GetTags(context.Background(), &svc.GetTagsRequest{Name: "ERC20"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"v1.0.0", "v1.0.1"}, tagsResp.GetTags())

	// Get MethodBySelector on default
	methodResp, err := r.GetMethodsBySelector(context.Background(),
		&svc.GetMethodsBySelectorRequest{
			Selector:        crypto.Keccak256(methodSig)[:4],
			AccountInstance: &common.AccountInstance{},
		})
	assert.NoError(t, err)
	assert.Nil(t, methodResp.GetMethod())
	assert.Equal(t, [][]byte{methodJSONs["isMinter"]}, methodResp.GetDefaultMethods())

	// Get EventsBySigHash wrong indexed count
	eventResp, err := r.GetEventsBySigHash(context.Background(),
		&svc.GetEventsBySigHashRequest{
			SigHash:           crypto.Keccak256Hash(eventSig).Bytes(),
			AccountInstance:   &ContractInstance,
			IndexedInputCount: 0})
	assert.Error(t, err)
	ierr, ok = err.(*ierror.Error)
	assert.True(t, ok, "GetEventsBySigHash error should cast to internal error")
	assert.Equal(t, "contract-registry.mock", ierr.GetComponent(), "GetEventsBySigHash error component should be correct")
	assert.True(t, errors.IsStorageError(ierr), "GetEventsBySigHash error should be a storage error")
	assert.Nil(t, eventResp.GetEvent())
	assert.Nil(t, eventResp.GetDefaultEvents())

	// Get EventsBySigHash
	eventResp, err = r.GetEventsBySigHash(context.Background(),
		&svc.GetEventsBySigHashRequest{
			SigHash:           crypto.Keccak256Hash(eventSig).Bytes(),
			AccountInstance:   &ContractInstance,
			IndexedInputCount: 1})
	assert.NoError(t, err)
	assert.Nil(t, eventResp.GetEvent())
	assert.Equal(t, [][]byte{eventJSONs["MinterAdded"], eventJSONsBis["MinterAdded"]}, eventResp.GetDefaultEvents())

	// Update smart-contract address
	_, err = r.SetAccountCodeHash(context.Background(),
		&svc.SetAccountCodeHashRequest{
			AccountInstance: &ContractInstance,
			CodeHash:        crypto.Keccak256([]byte{1, 2, 3}),
		})
	assert.NoError(t, err)

	// Get MethodBySelector
	methodResp, err = r.GetMethodsBySelector(context.Background(),
		&svc.GetMethodsBySelectorRequest{
			Selector:        crypto.Keccak256(methodSig)[:4],
			AccountInstance: &ContractInstance})
	assert.NoError(t, err)
	assert.Equal(t, methodJSONs["isMinter"], methodResp.GetMethod())
	assert.Nil(t, methodResp.GetDefaultMethods())

	// Get EventsBySigHash
	eventResp, err = r.GetEventsBySigHash(
		context.Background(),
		&svc.GetEventsBySigHashRequest{
			SigHash:           crypto.Keccak256Hash(eventSig).Bytes(),
			AccountInstance:   &ContractInstance,
			IndexedInputCount: 1})
	assert.NoError(t, err)
	assert.Equal(t, eventJSONs["MinterAdded"], eventResp.GetEvent())
	assert.Nil(t, eventResp.GetDefaultEvents())
}
