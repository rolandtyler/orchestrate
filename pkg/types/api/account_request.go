package api

import (
	qkmtypes "github.com/ConsenSys/orchestrate/pkg/quorum-key-manager/types"
)

type CreateAccountRequest struct {
	Alias      string            `json:"alias" validate:"omitempty" example:"personal-account"`
	Chain      string            `json:"chain" validate:"omitempty" example:"besu"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type ImportAccountRequest struct {
	Alias      string            `json:"alias" validate:"omitempty" example:"personal-account"`
	Chain      string            `json:"chain" validate:"omitempty" example:"quorum"`
	PrivateKey string            `json:"privateKey" validate:"required" example:"66232652FDFFD802B7252A456DBD8F3ECC0352BBDE76C23B40AFE8AEBD714E2D"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type UpdateAccountRequest struct {
	Alias      string            `json:"alias" validate:"omitempty"  example:"personal-account"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type SignPayloadRequest struct {
	Data string `json:"data" validate:"required,isHex" example:"0x6d79206461746120746f207369676e"`
}

type SignTypedDataRequest struct {
	DomainSeparator qkmtypes.DomainSeparator   `json:"domainSeparator" validate:"required"`
	Types           map[string][]qkmtypes.Type `json:"types" validate:"required"`
	Message         map[string]interface{}     `json:"message" validate:"required"`
	MessageType     string                     `json:"messageType" validate:"required" example:"Mail"`
}
