package types

import (
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
)

type Annotations struct {
	OneTimeKey bool                 `json:"oneTimeKey,omitempty" example:"true"`
	ChainID    string               `json:"chainID,omitempty" example:"1 (mainnet)"`
	Priority   string               `json:"priority,omitempty" validate:"isPriority" example:"very-high"`
}

func (a *Annotations) Validate() error {
	if a == nil {
		return nil
	}

	if err := utils.GetValidator().Struct(a); err != nil {
		return err
	}

	return nil
}
