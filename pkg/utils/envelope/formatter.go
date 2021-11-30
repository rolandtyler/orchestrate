package envelope

import (
	"context"

	authutils "github.com/consensys/orchestrate/pkg/toolkit/app/auth/utils"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/types/ethereum"
	"github.com/consensys/orchestrate/pkg/types/tx"
	"github.com/consensys/orchestrate/pkg/utils"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewEnvelopeFromJob(job *entities.Job, headers map[string]string) *tx.TxEnvelope {
	contextLabels := job.Labels
	if contextLabels == nil {
		contextLabels = map[string]string{}
	}

	contextLabels[tx.NextJobUUIDLabel] = job.NextJobUUID
	contextLabels[tx.PriorityLabel] = job.InternalData.Priority
	contextLabels[tx.ParentJobUUIDLabel] = job.InternalData.ParentJobUUID

	txEnvelope := &tx.TxEnvelope{
		Msg: &tx.TxEnvelope_TxRequest{TxRequest: &tx.TxRequest{
			Id:      job.ScheduleUUID,
			Headers: headers,
			Params: &tx.Params{
				From:            utils.HexToString(job.Transaction.From),
				To:              utils.HexToString(job.Transaction.To),
				Gas:             utils.IntegerToString(job.Transaction.Gas),
				GasPrice:        utils.HexToString(job.Transaction.GasPrice),
				GasFeeCap:       utils.HexToString(job.Transaction.GasFeeCap),
				GasTipCap:       utils.HexToString(job.Transaction.GasTipCap),
				Value:           utils.HexToString(job.Transaction.Value),
				Nonce:           utils.IntegerToString(job.Transaction.Nonce),
				Data:            utils.HexToString(job.Transaction.Data),
				Raw:             utils.HexToString(job.Transaction.Raw),
				PrivateFrom:     utils.BytesToString(job.Transaction.PrivateFrom),
				PrivateFor:      utils.ArrBytesToString(job.Transaction.PrivateFor),
				MandatoryFor:    utils.ArrBytesToString(job.Transaction.MandatoryFor),
				PrivacyGroupId:  utils.BytesToString(job.Transaction.PrivacyGroupID),
				PrivacyFlag:     int32(job.Transaction.PrivacyFlag),
				TransactionType: string(job.Transaction.TransactionType),
				AccessList:      ConvertFromAccessList(job.Transaction.AccessList),
			},
			ContextLabels: contextLabels,
			JobType:       tx.JobTypeMap[job.Type],
		}},
		InternalLabels: make(map[string]string),
	}

	txEnvelope.SetChainUUID(job.ChainUUID)

	txEnvelope.SetChainID(job.InternalData.ChainID)
	txEnvelope.SetScheduleUUID(job.ScheduleUUID)
	txEnvelope.SetJobUUID(job.UUID)

	if job.InternalData.OneTimeKey {
		txEnvelope.EnableTxFromOneTimeKey()
	}

	if job.InternalData.ParentJobUUID != "" {
		txEnvelope.SetParentJobUUID(job.InternalData.ParentJobUUID)
	}

	if job.InternalData.Priority != "" {
		txEnvelope.SetPriority(job.InternalData.Priority)
	}

	if job.Transaction.Hash != nil {
		txEnvelope.SetTxHash(job.Transaction.Hash.String())
	}

	return txEnvelope
}

func NewContextFromEnvelope(ctx context.Context, envelope *tx.Envelope) context.Context {
	return multitenancy.WithUserInfo(ctx, multitenancy.NewUserInfo(
		envelope.GetHeadersValue(authutils.TenantIDHeader),
		envelope.GetHeadersValue(authutils.UsernameHeader),
	))
}

func NewJobFromEnvelope(envelope *tx.Envelope) *entities.Job {
	return &entities.Job{
		UUID:         envelope.GetJobUUID(),
		NextJobUUID:  envelope.GetNextJobUUID(),
		ChainUUID:    envelope.GetChainUUID(),
		ScheduleUUID: envelope.GetScheduleUUID(),
		Type:         entities.JobType(envelope.GetJobTypeString()),
		InternalData: &entities.InternalData{
			OneTimeKey:    envelope.IsOneTimeKeySignature(),
			ChainID:       envelope.GetChainID(),
			ParentJobUUID: envelope.GetParentJobUUID(),
			Priority:      envelope.GetPriority(),
		},
		TenantID: envelope.GetHeadersValue(authutils.TenantIDHeader),
		OwnerID:  envelope.GetHeadersValue(authutils.UsernameHeader),
		Transaction: &entities.ETHTransaction{
			Hash:            envelope.GetTxHash(),
			From:            envelope.GetFrom(),
			To:              envelope.GetTo(),
			Nonce:           envelope.GetNonce(),
			Value:           (*hexutil.Big)(envelope.GetValue()),
			GasPrice:        (*hexutil.Big)(envelope.GetGasPrice()),
			Gas:             envelope.GetGas(),
			GasFeeCap:       (*hexutil.Big)(envelope.GetGasFeeCap()),
			GasTipCap:       (*hexutil.Big)(envelope.GetGasTipCap()),
			AccessList:      ConvertToAccessList(envelope.GetAccessList()),
			TransactionType: entities.TransactionType(envelope.GetTransactionType()),
			Data:            envelope.GetData(),
			Raw:             envelope.GetRaw(),
			PrivateFrom:     []byte(envelope.GetPrivateFrom()),
			PrivateFor:      utils.ArrStringToBytes(envelope.GetPrivateFor()),
			MandatoryFor:    utils.ArrStringToBytes(envelope.GetMandatoryFor()),
			PrivacyGroupID:  []byte(envelope.GetPrivacyGroupID()),
			PrivacyFlag:     envelope.GetPrivacyFlag(),
			EnclaveKey:      envelope.GetEnclaveKey(),
		},
	}
}

func ConvertFromAccessList(accessList types.AccessList) []*ethereum.AccessTuple {
	result := []*ethereum.AccessTuple{}
	for _, t := range accessList {
		tupl := &ethereum.AccessTuple{
			Address:     t.Address.Hex(),
			StorageKeys: []string{},
		}

		for _, k := range t.StorageKeys {
			tupl.StorageKeys = append(tupl.StorageKeys, k.Hex())
		}

		result = append(result, tupl)
	}

	return result
}

func ConvertToAccessList(accessList []*ethereum.AccessTuple) types.AccessList {
	result := types.AccessList{}
	for _, item := range accessList {
		storageKeys := []ethcommon.Hash{}
		for _, sk := range item.StorageKeys {
			storageKeys = append(storageKeys, ethcommon.HexToHash(sk))
		}

		result = append(result, types.AccessTuple{
			Address:     ethcommon.HexToAddress(item.Address),
			StorageKeys: storageKeys,
		})
	}

	return result
}
