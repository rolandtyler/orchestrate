package parsers

import (
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/consensys/orchestrate/services/api/store/models"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewTransactionModelFromEntities(tx *entities.ETHTransaction) *models.Transaction {
	return &models.Transaction{
		Hash:           utils.StructToString(tx.Hash),
		Sender:         utils.StructToString(tx.From),
		Recipient:      utils.StructToString(tx.To),
		Nonce:          utils.ValueToString(tx.Nonce),
		Value:          utils.StructToString(tx.Value),
		GasPrice:       utils.StructToString(tx.GasPrice),
		GasFeeCap:      utils.StructToString(tx.GasFeeCap),
		GasTipCap:      utils.StructToString(tx.GasTipCap),
		Gas:            utils.ValueToString(tx.Gas),
		Data:           utils.StructToString(tx.Data),
		Raw:            utils.StructToString(tx.Raw),
		TxType:         string(tx.TransactionType),
		AccessList:     tx.AccessList,
		PrivateFrom:    utils.BytesToString(tx.PrivateFrom),
		PrivateFor:     utils.ArrBytesToString(tx.PrivateFor),
		MandatoryFor:   utils.ArrBytesToString(tx.MandatoryFor),
		PrivacyGroupID: utils.BytesToString(tx.PrivacyGroupID),
		PrivacyFlag:    int(tx.PrivacyFlag),
		EnclaveKey:     utils.StructToString(tx.EnclaveKey),
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
		PrivateFrom:     []byte(tx.PrivateFrom),
		PrivateFor:      utils.ArrStringToBytes(tx.PrivateFor),
		MandatoryFor:    utils.ArrStringToBytes(tx.MandatoryFor),
		PrivacyGroupID:  []byte(tx.PrivacyGroupID),
		PrivacyFlag:     entities.PrivacyFlag(tx.PrivacyFlag),
		EnclaveKey:      utils.StringToHexBytes(tx.EnclaveKey),
		Raw:             utils.StringToHexBytes(tx.Raw),
		CreatedAt:       tx.CreatedAt,
		UpdatedAt:       tx.UpdatedAt,
	}
}

func UpdateTransactionModelFromEntities(txModel *models.Transaction, tx *entities.ETHTransaction) {
	txModel.Hash = utils.StructToString(tx.Hash)
	txModel.Sender = utils.StructToString(tx.From)
	txModel.Recipient = utils.StructToString(tx.To)
	txModel.Nonce = utils.ValueToString(tx.Nonce)
	txModel.Value = utils.StructToString(tx.Value)
	txModel.GasPrice = utils.StructToString(tx.GasPrice)
	txModel.GasFeeCap = utils.StructToString(tx.GasFeeCap)
	txModel.GasTipCap = utils.StructToString(tx.GasTipCap)
	txModel.Gas = utils.ValueToString(tx.Gas)
	txModel.Data = utils.StructToString(tx.Data)
	txModel.Raw = utils.StructToString(tx.Raw)
	txModel.TxType = string(tx.TransactionType)
	txModel.AccessList = tx.AccessList
	txModel.PrivateFrom = utils.BytesToString(tx.PrivateFrom)
	txModel.PrivateFor = utils.ArrBytesToString(tx.PrivateFor)
	txModel.MandatoryFor = utils.ArrBytesToString(tx.MandatoryFor)
	txModel.PrivacyGroupID = utils.BytesToString(tx.PrivacyGroupID)
	txModel.EnclaveKey = utils.StructToString(tx.EnclaveKey)
}
