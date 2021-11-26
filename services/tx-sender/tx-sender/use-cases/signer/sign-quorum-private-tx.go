package signer

import (
	"context"

	pkgcryto "github.com/consensys/orchestrate/pkg/crypto/ethereum"
	"github.com/consensys/orchestrate/pkg/encoding/rlp"
	qkm "github.com/consensys/orchestrate/pkg/quorum-key-manager"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/services/tx-sender/tx-sender/parsers"
	qkmtypes "github.com/consensys/quorum-key-manager/src/stores/api/types"
	quorumtypes "github.com/consensys/quorum/core/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	usecases "github.com/consensys/orchestrate/services/tx-sender/tx-sender/use-cases"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/quorum-key-manager/pkg/client"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const signQuorumPrivateTransactionComponent = "use-cases.sign-quorum-private-transaction"

// signQuorumPrivateTransactionUseCase is a use case to sign a quorum private transaction
type signQuorumPrivateTransactionUseCase struct {
	keyManagerClient client.KeyManagerClient
	logger           *log.Logger
	storeName        string
}

// NewSignQuorumPrivateTransactionUseCase creates a new signQuorumPrivateTransactionUseCase
func NewSignQuorumPrivateTransactionUseCase(keyManagerClient client.KeyManagerClient) usecases.SignQuorumPrivateTransactionUseCase {
	return &signQuorumPrivateTransactionUseCase{
		keyManagerClient: keyManagerClient,
		logger:           log.NewLogger().SetComponent(signQuorumPrivateTransactionComponent),
		storeName:        qkm.GlobalStoreName(),
	}
}

// Execute signs a quorum private transaction
func (uc *signQuorumPrivateTransactionUseCase) Execute(ctx context.Context, job *entities.Job) (signedRaw hexutil.Bytes, txHash ethcommon.Hash, err error) {
	logger := uc.logger.WithContext(ctx).WithField("one_time_key", job.InternalData.OneTimeKey)

	transaction := parsers.ETHTransactionToQuorumTransaction(job.Transaction)
	transaction.SetPrivate()
	if job.InternalData.OneTimeKey {
		signedRaw, txHash, err = uc.signWithOneTimeKey(ctx, transaction)
	} else {
		signedRaw, txHash, err = uc.signWithAccount(ctx, job, transaction)
	}
	if err != nil {
		return nil, ethcommon.Hash{}, errors.FromError(err).ExtendComponent(signQuorumPrivateTransactionComponent)
	}

	logger.WithField("tx_hash", txHash).Debug("quorum private transaction signed successfully")
	return signedRaw, txHash, nil
}

func (uc *signQuorumPrivateTransactionUseCase) signWithOneTimeKey(ctx context.Context, transaction *quorumtypes.Transaction) (signedRaw hexutil.Bytes, txHash ethcommon.Hash, err error) {
	logger := uc.logger.WithContext(ctx)
	privKey, err := crypto.GenerateKey()
	if err != nil {
		errMessage := "failed to generate Ethereum account"
		logger.WithError(err).Error(errMessage)
		return nil, ethcommon.Hash{}, errors.CryptoOperationError(errMessage)
	}

	signer := pkgcryto.GetQuorumPrivateTxSigner()
	decodedSignature, err := pkgcryto.SignQuorumPrivateTransaction(transaction, privKey, signer)
	if err != nil {
		logger.WithError(err).Error("failed to sign private transaction")
		return nil, ethcommon.Hash{}, err
	}

	signedTx, err := transaction.WithSignature(signer, decodedSignature)
	if err != nil {
		errMessage := "failed to set quorum private transaction signature"
		logger.WithError(err).Error(errMessage)
		return nil, ethcommon.Hash{}, errors.InvalidParameterError(errMessage).ExtendComponent(signQuorumPrivateTransactionComponent)
	}

	signedRawB, err := rlp.Encode(signedTx)
	if err != nil {
		errMessage := "failed to RLP encode signed quorum private transaction"
		logger.WithError(err).Error(errMessage)
		return nil, ethcommon.Hash{}, errors.CryptoOperationError(errMessage).ExtendComponent(signQuorumPrivateTransactionComponent)
	}

	return signedRawB, signedTx.Hash(), nil
}

func (uc *signQuorumPrivateTransactionUseCase) signWithAccount(ctx context.Context, job *entities.Job, tx *quorumtypes.Transaction) (signedRaw hexutil.Bytes, txHash ethcommon.Hash, err error) {
	logger := uc.logger.WithContext(ctx)
	signedRawStr, err := uc.keyManagerClient.SignQuorumPrivateTransaction(ctx, uc.storeName, job.Transaction.From.Hex(), &qkmtypes.SignQuorumPrivateTransactionRequest{
		Nonce:    hexutil.Uint64(tx.Nonce()),
		To:       tx.To(),
		Value:    hexutil.Big(*tx.Value()),
		GasPrice: hexutil.Big(*tx.GasPrice()),
		GasLimit: hexutil.Uint64(tx.Gas()),
		Data:     tx.Data(),
	})
	if err != nil {
		errMsg := "failed to sign quorum private transaction using key manager"
		logger.WithError(err).Error(errMsg)
		return nil, ethcommon.Hash{}, errors.DependencyFailureError(errMsg).AppendReason(err.Error())
	}

	signedRaw, err = hexutil.Decode(signedRawStr)
	if err != nil {
		errMessage := "failed to decode signature"
		logger.WithError(err).Error(errMessage)
		return nil, ethcommon.Hash{}, errors.EncodingError(errMessage)
	}

	err = rlp.Decode(signedRaw, &tx)
	if err != nil {
		errMessage := "failed to decode quorum signed transaction"
		logger.WithError(err).Error(errMessage)
		return nil, ethcommon.Hash{}, errors.EncodingError(errMessage)
	}

	return signedRaw, tx.Hash(), nil
}
