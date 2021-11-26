package parsers

import (
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/consensys/orchestrate/services/api/store/models"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewTransactionModelFromEntities(tx *entities.ETHTransaction) *models.Transaction {
	return &models.Transaction{
		Hash:           utils.ObjectToString(tx.Hash),
		Sender:         utils.ObjectToString(tx.From),
		Recipient:      utils.ObjectToString(tx.To),
		Nonce:          utils.ObjectToString(tx.Nonce),
		Value:          utils.ObjectToString(tx.Value),
		GasPrice:       utils.ObjectToString(tx.GasPrice),
		GasFeeCap:      utils.ObjectToString(tx.GasFeeCap),
		GasTipCap:      utils.ObjectToString(tx.GasTipCap),
		Gas:            utils.ObjectToString(tx.Gas),
		Data:           utils.ObjectToString(tx.Data),
		Raw:            utils.ObjectToString(tx.Raw),
		TxType:         string(tx.TransactionType),
		AccessList:     tx.AccessList,
		PrivateFrom:    utils.BytesToString(tx.PrivateFrom),
		PrivateFor:     utils.ArrBytesToString(tx.PrivateFor),
		MandatoryFor:   utils.ArrBytesToString(tx.MandatoryFor),
		PrivacyGroupID: utils.BytesToString(tx.PrivacyGroupID),
		PrivacyFlag:    int(tx.PrivacyFlag),
		EnclaveKey:     utils.ObjectToString(tx.EnclaveKey),
		CreatedAt:      tx.CreatedAt,
		UpdatedAt:      tx.UpdatedAt,
	}
}

func NewTransactionEntityFromModels(tx *models.Transaction) *entities.ETHTransaction {
	accessList := types.AccessList{}
	_ = utils.CastInterfaceToObject(tx.AccessList, &accessList)

	return &entities.ETHTransaction{
		Hash:            utils.StringToEthHash(tx.Hash),
		From:            utils.ToEthAddr(tx.Sender),
		To:              utils.ToEthAddr(tx.Recipient),
		Nonce:           tx.Nonce,
		Value:           tx.Value,
		GasPrice:        tx.GasPrice,
		Gas:             tx.Gas,
		GasTipCap:       tx.GasTipCap,
		GasFeeCap:       tx.GasFeeCap,
		Data:            tx.Data,
		TransactionType: entities.TransactionType(tx.TxType),
		AccessList:      accessList,
		PrivateFrom:     tx.PrivateFrom,
		PrivateFor:      tx.PrivateFor,
		MandatoryFor:    tx.MandatoryFor,
		PrivacyGroupID:  tx.PrivacyGroupID,
		PrivacyFlag:     entities.PrivacyFlag(tx.PrivacyFlag),
		EnclaveKey:      tx.EnclaveKey,
		Raw:             tx.Raw,
		CreatedAt:       tx.CreatedAt,
		UpdatedAt:       tx.UpdatedAt,
	}
}

func UpdateTransactionModelFromEntities(txModel *models.Transaction, tx *entities.ETHTransaction) {
	txModel.Hash = utils.ObjectToString(tx.Hash)
	txModel.Sender = utils.ObjectToString(tx.From)
	txModel.Recipient = utils.ObjectToString(tx.To)
	txModel.Nonce = utils.ObjectToString(tx.Nonce)
	txModel.Value = utils.ObjectToString(tx.Value)
	txModel.GasPrice = utils.ObjectToString(tx.GasPrice)
	txModel.GasFeeCap = utils.ObjectToString(tx.GasFeeCap)
	txModel.GasTipCap = utils.ObjectToString(tx.GasTipCap)
	txModel.Gas = utils.ObjectToString(tx.Gas)
	txModel.Data = utils.ObjectToString(tx.Data)
	txModel.Raw = utils.ObjectToString(tx.Raw)
	if tx.TransactionType != "" {
		txModel.TxType = string(tx.TransactionType)
	}
	txModel.AccessList = tx.AccessList
	txModel.PrivateFrom = utils.BytesToString(tx.PrivateFrom)
	txModel.PrivateFor = utils.ArrBytesToString(tx.PrivateFor)
	txModel.MandatoryFor = utils.ArrBytesToString(tx.MandatoryFor)
	txModel.PrivacyGroupID = utils.BytesToString(tx.PrivacyGroupID)
	txModel.EnclaveKey = utils.ObjectToString(tx.EnclaveKey)
}
