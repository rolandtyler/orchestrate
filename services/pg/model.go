package pg

import (
	"fmt"
	"strings"
	"time"

	encoding "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/encoding/proto"
	evlpstore "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/services/envelope-store"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/envelope"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/utils"
)

// EnvelopeModel represent elements in `envelopes` table
type EnvelopeModel struct {
	tableName struct{} `sql:"envelopes"` //nolint:unused,structcheck

	// ID technical identifier
	ID int32

	// Envelope Identifier
	EnvelopeID string

	// Ethereum info about transaction
	ChainID string
	TxHash  string

	// Envelope
	Envelope []byte

	// Status
	Status   string
	StoredAt time.Time
	SentAt   time.Time
	MinedAt  time.Time
	ErrorAt  time.Time
}

// StatusInfo returns a proto formated StatusInfo object from a model
func (model *EnvelopeModel) StatusInfo() *evlpstore.StatusInfo {
	return &evlpstore.StatusInfo{
		Status:   model.StatusFormated(),
		StoredAt: utils.TimeToPTimestamp(model.StoredAt),
		SentAt:   utils.TimeToPTimestamp(model.SentAt),
		MinedAt:  utils.TimeToPTimestamp(model.MinedAt),
		ErrorAt:  utils.TimeToPTimestamp(model.ErrorAt),
	}
}

// StatusFormated returns a proto Status enum
func (model *EnvelopeModel) StatusFormated() evlpstore.Status {
	status, ok := evlpstore.Status_value[strings.ToUpper(model.Status)]
	if !ok {
		panic(fmt.Sprintf("invalid status %q", model.Status))
	}
	return evlpstore.Status(status)
}

// UnmarshalEnvelope returns a proto envelope from model
func (model *EnvelopeModel) UnmarshalEnvelope(e *envelope.Envelope) error {
	return encoding.Unmarshal(model.Envelope, e)
}

func (model *EnvelopeModel) ToStatusResponse() (*evlpstore.StatusResponse, error) {
	return &evlpstore.StatusResponse{
		StatusInfo: model.StatusInfo(),
	}, nil
}

func (model *EnvelopeModel) ToStoreResponse() (*evlpstore.StoreResponse, error) {
	resp := &evlpstore.StoreResponse{
		StatusInfo: model.StatusInfo(),
		Envelope:   &envelope.Envelope{},
	}

	// Unmarshal envelope
	err := model.UnmarshalEnvelope(resp.GetEnvelope())
	if err != nil {
		return &evlpstore.StoreResponse{}, err
	}

	return resp, nil
}

// FromEnvelope creates a model from an envelope
func FromEnvelope(e *envelope.Envelope) (*EnvelopeModel, error) {
	// Marshal envelope
	b, err := encoding.Marshal(e)
	if err != nil {
		return nil, err
	}

	return &EnvelopeModel{
		Envelope:   b,
		EnvelopeID: e.GetMetadata().GetId(),
		ChainID:    e.GetChain().ID().String(),
		TxHash:     e.GetTx().GetHash().Hex(),
	}, nil
}
