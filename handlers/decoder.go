package handlers

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/core.git/services"
	"gitlab.com/ConsenSys/client/fr/core-stack/core.git/types"
	InfEth "gitlab.com/ConsenSys/client/fr/core-stack/infra/ethereum.git"
)

// LogDecoder decode a single log
func LogDecoder(ctx *types.Context, r services.ABIRegistry, log *types.Log, i int) {
	event, err := r.GetEventBySig(log.Topics[0].Hex())
	if err != nil {
		ctx.Error(err)
		return
	}

	mapping, err := InfEth.Decode(&event, &log.Log)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Logger.WithFields(logrus.Fields{
		"DecodedData": mapping,
	}).Debug("Event decoded")
	ctx.T.Receipt().Logs[i].SetDecodedData(mapping)

}

// TransactionDecoder creates a decode handler
func TransactionDecoder(r services.ABIRegistry) types.HandlerFunc {
	return func(ctx *types.Context) {

		for i, log := range ctx.T.Receipt().Logs {
			go LogDecoder(ctx, r, log, i)
		}

		return
	}
}
