package api

import (
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
)

type RegisterContractRequest struct {
	ABI              interface{} `json:"abi,omitempty" validate:"required" example:"[{anonymous: false, inputs: [{indexed: false, name: account, type: address}, name: MinterAdded, type: event}]}]"`
	Bytecode         string      `json:"bytecode,omitempty" validate:"omitempty,isHex" example:"0x6080604052348015600f57600080f..."`
	DeployedBytecode string      `json:"deployed_byte_code,omitempty" validate:"omitempty,isHex" example:"0x6080604052348015600f57600080f..."`
	Name             string      `json:"name" validate:"required" example:"ERC20"`
	Tag              string      `json:"tag,omitempty" example:"v1.0.0"`
}

type ContractResponse struct {
	*entities.Contract
}

type GetContractEventsBySignHashRequest struct {
	ChainID           string `json:"chain_id" validate:"required" example:"2017"`
	SigHash           string `json:"sig_hash" validate:"required,isHex" example:"0x6080604052348015600f57600080f..."`
	IndexedInputCount uint32 `json:"indexed_input_count" validate:"omitempty" example:"1"`
}

type GetContractEventsBySignHashResponse struct {
	Event         string   `json:"event" validate:"omitempty" example:"{anonymous:false,inputs:[{indexed:true,name:from,type:address},{indexed:true,name:to,type:address},{indexed:false,name:value,type:uint256}],name:Transfer,type:event}"`
	DefaultEvents []string `json:"default_events" validate:"omitempty" example:"[{anonymous:false,inputs:[{indexed:true,name:from,type:address},{indexed:true,name:to,type:address},{indexed:false,name:value,type:uint256}],name:Transfer,type:event},..."`
}

type SetContractCodeHashRequest struct {
	Address  string `json:"address" validate:"required,isHexAddress" example:"0xca35b7d915458ef540ade6068dfe2f44e8fa733c"`
	ChainID  string `json:"chain_id" validate:"required" example:"2017"`
	CodeHash string `json:"code_hash" validate:"required,isHex" example:"0x6080604052348015600f57600080f..."`
}