package api

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type RegisterFaucetRequest struct {
	Name            string            `json:"name" validate:"required" example:"faucet-mainnet"`
	ChainRule       string            `json:"chainRule" validate:"required" example:"mainnet"`
	CreditorAccount ethcommon.Address `json:"creditorAccount" validate:"required" example:"0x905B88EFf8Bda1543d4d6f4aA05afef143D27E18" swaggertype:"string"`
	MaxBalance      string            `json:"maxBalance" validate:"required,isBig" example:"100000000000000000 (wei)"`
	Amount          string            `json:"amount" validate:"required,isBig" example:"60000000000000000 (wei)"`
	Cooldown        string            `json:"cooldown" validate:"required,isDuration" example:"10s"`
}

type UpdateFaucetRequest struct {
	Name            string            `json:"name,omitempty" validate:"omitempty" example:"faucet-mainnet"`
	ChainRule       string            `json:"chainRule,omitempty" validate:"omitempty" example:"mainnet"`
	CreditorAccount ethcommon.Address `json:"creditorAccount,omitempty" validate:"omitempty" example:"0x905B88EFf8Bda1543d4d6f4aA05afef143D27E18" swaggertype:"string"`
	MaxBalance      string            `json:"maxBalance,omitempty" validate:"omitempty,isBig" example:"100000000000000000 (wei)"`
	Amount          string            `json:"amount,omitempty" validate:"omitempty,isBig" example:"60000000000000000 (wei)"`
	Cooldown        string            `json:"cooldown,omitempty" validate:"omitempty,isDuration" example:"10s"`
}
