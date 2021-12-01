package parsers

import (
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/consensys/orchestrate/services/api/store/models"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewTransactionModelFromEntities(tx *entities.ETHTransaction) *models.Transaction {
	return &models.Transaction{
		Hash:           utils.ObjToString(tx.Hash),
		Sender:         utils.ObjToString(tx.From),
		Recipient:      utils.ObjToString(tx.To),
		Nonce:          utils.ValueToString(tx.Nonce),
		Value:          utils.ObjToString(tx.Value),
		GasPrice:       utils.ObjToString(tx.GasPrice),
		GasFeeCap:      utils.ObjToString(tx.GasFeeCap),
		GasTipCap:      utils.ObjToString(tx.GasTipCap),
		Gas:            utils.ValueToString(tx.Gas),
		Data:           utils.ObjToString(tx.Data),
		Raw:            utils.ObjToString(tx.Raw),
		TxType:         string(tx.TransactionType),
		AccessList:     tx.AccessList,
		PrivateFrom:    tx.PrivateFrom,
		PrivateFor:     tx.PrivateFor,
		MandatoryFor:   tx.MandatoryFor,
		PrivacyGroupID: tx.PrivacyGroupID,
		PrivacyFlag:    int(tx.PrivacyFlag),
		EnclaveKey:     utils.ObjToString(tx.EnclaveKey),
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
		Nonce:           utils.StringToUint64(tx.Nonce),
		Value:           utils.StringToHexInt(tx.Value),
		GasPrice:        utils.StringToHexInt(tx.GasPrice),
		Gas:             utils.StringToUint64(tx.Gas),
		GasTipCap:       utils.StringToHexInt(tx.GasTipCap),
		GasFeeCap:       utils.StringToHexInt(tx.GasFeeCap),
		Data:            utils.StringToHexBytes(tx.Data),
		TransactionType: entities.TransactionType(tx.TxType),
		AccessList:      accessList,
		PrivateFrom:     tx.PrivateFrom,
		PrivateFor:      tx.PrivateFor,
		MandatoryFor:    tx.MandatoryFor,
		PrivacyGroupID:  tx.PrivacyGroupID,
		PrivacyFlag:     entities.PrivacyFlag(tx.PrivacyFlag),
		EnclaveKey:      utils.StringToHexBytes(tx.EnclaveKey),
		Raw:             utils.StringToHexBytes(tx.Raw),
		CreatedAt:       tx.CreatedAt,
		UpdatedAt:       tx.UpdatedAt,
	}
}

func UpdateTransactionModelFromEntities(txModel *models.Transaction, tx *entities.ETHTransaction) {
	txModel.Hash = utils.ObjToString(tx.Hash)
	txModel.Sender = utils.ObjToString(tx.From)
	txModel.Recipient = utils.ObjToString(tx.To)
	txModel.Nonce = utils.ValueToString(tx.Nonce)
	txModel.Value = utils.ObjToString(tx.Value)
	txModel.GasPrice = utils.ObjToString(tx.GasPrice)
	txModel.GasFeeCap = utils.ObjToString(tx.GasFeeCap)
	txModel.GasTipCap = utils.ObjToString(tx.GasTipCap)
	txModel.Gas = utils.ValueToString(tx.Gas)
	txModel.Data = utils.ObjToString(tx.Data)
	txModel.Raw = utils.ObjToString(tx.Raw)
	txModel.TxType = string(tx.TransactionType)
	txModel.AccessList = tx.AccessList
	txModel.PrivateFrom = tx.PrivateFrom
	txModel.PrivateFor = tx.PrivateFor
	txModel.MandatoryFor = tx.MandatoryFor
	txModel.PrivacyGroupID = tx.PrivacyGroupID
	txModel.EnclaveKey = utils.ObjToString(tx.EnclaveKey)
}
