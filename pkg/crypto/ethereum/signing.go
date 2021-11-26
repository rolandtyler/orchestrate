package ethereum

import (
	"crypto/ecdsa"
	"math/big"

	quorumtypes "github.com/consensys/quorum/core/types"

	"github.com/consensys/orchestrate/pkg/encoding/rlp"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignTransaction(tx *types.Transaction, privKey *ecdsa.PrivateKey, signer types.Signer) ([]byte, error) {
	h := signer.Hash(tx)
	decodedSignature, err := crypto.Sign(h[:], privKey)
	if err != nil {
		return nil, errors.CryptoOperationError(err.Error())
	}

	return decodedSignature, nil
}

func SignQuorumPrivateTransaction(tx *quorumtypes.Transaction, privKey *ecdsa.PrivateKey, signer quorumtypes.Signer) ([]byte, error) {
	h := signer.Hash(tx)
	decodedSignature, err := crypto.Sign(h[:], privKey)
	if err != nil {
		return nil, errors.CryptoOperationError(err.Error())
	}

	return decodedSignature, nil
}

func SignEEATransaction(tx *types.Transaction, privateArgs *entities.PrivateETHTransactionParams, chainID *big.Int, privKey *ecdsa.PrivateKey) ([]byte, error) {
	hash, err := EEATransactionPayload(tx, privateArgs, chainID)
	if err != nil {
		return nil, err
	}

	signature, err := crypto.Sign(hash, privKey)
	if err != nil {
		return nil, errors.CryptoOperationError("failed to sign eea transaction").AppendReason(err.Error())
	}

	return signature, err
}

func EEATransactionPayload(tx *types.Transaction, privateArgs *entities.PrivateETHTransactionParams, chainID *big.Int) ([]byte, error) {
	privateFromEncoded, err := GetEncodedPrivateFrom(privateArgs.PrivateFrom)
	if err != nil {
		return nil, err
	}

	privateRecipientEncoded, err := GetEncodedPrivateRecipient(privateArgs.PrivacyGroupID, privateArgs.PrivateFor)
	if err != nil {
		return nil, err
	}

	hash, err := rlp.Hash([]interface{}{
		tx.Nonce(),
		tx.GasPrice(),
		tx.Gas(),
		tx.To(),
		tx.Value(),
		tx.Data(),
		chainID,
		uint(0),
		uint(0),
		privateFromEncoded,
		privateRecipientEncoded,
		privateArgs.PrivateTxType,
	})
	if err != nil {
		return nil, errors.CryptoOperationError("failed to hash eea transaction").AppendReason(err.Error())
	}

	return hash.Bytes(), nil
}
