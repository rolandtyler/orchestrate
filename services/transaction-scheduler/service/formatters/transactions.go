package formatters

import (
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/service/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/entities"
)

func FormatSendTxRequest(txRequest *types.SendTransactionRequest, chainUUID string) *entities.TxRequest {
	return &entities.TxRequest{
		IdempotencyKey: txRequest.IdempotencyKey,
		Labels:         txRequest.Labels,
		Schedule: &entities.Schedule{
			ChainUUID: chainUUID,
		},
		Params: &entities.TxRequestParams{
			From:            txRequest.Params.From,
			To:              txRequest.Params.To,
			Value:           txRequest.Params.Value,
			GasPrice:        txRequest.Params.GasPrice,
			GasLimit:        txRequest.Params.Gas,
			MethodSignature: txRequest.Params.MethodSignature,
			Args:            txRequest.Params.Args,
		},
	}
}

func FormatTxResponse(txRequest *entities.TxRequest) *types.TransactionResponse {
	return &types.TransactionResponse{
		IdempotencyKey: txRequest.IdempotencyKey,
		Params:         txRequest.Params,
		Schedule:       FormatScheduleResponse(txRequest.Schedule),
		CreatedAt:      txRequest.CreatedAt,
	}
}