package sender

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/errors"
	evlpstore "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/services/envelope-store"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/ethereum.git/ethclient"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-sender.git/handlers/store"
)

// Sender creates sender handler
func Sender(ec ethclient.TransactionSender, s evlpstore.EnvelopeStoreClient) engine.HandlerFunc {
	// Declare a set of handlers that will be forked by Sender handler
	rawTxStore := store.RawTxStore(s)
	UnsignedTxStore := store.UnsignedTxStore(s)

	rawTxSender := engine.CombineHandlers(
		rawTxStore,
		RawTxSender(ec),
	)

	rawPrivateTxSender := engine.CombineHandlers(
		UnsignedTxStore,
		RawPrivateTxSender(ec),
	)

	tesseraRawPrivateTxSender := engine.CombineHandlers(
		rawTxStore,
		TesseraRawPrivateTxSender(ec),
	)

	unsignedTxSender := engine.CombineHandlers(
		UnsignedTxStore,
		UnsignedTxSender(ec),
	)

	return func(txctx *engine.TxContext) {
		txctx.Logger = txctx.Logger.WithFields(log.Fields{
			"chain.id":    txctx.Envelope.GetChain().ID().String(),
			"metadata.id": txctx.Envelope.GetMetadata().GetId(),
			"tx.raw":      utils.ShortString(txctx.Envelope.GetTx().GetRaw().Hex(), 30),
			"tx.hash":     txctx.Envelope.GetTx().GetHash().Hex(),
			"from":        txctx.Envelope.GetFrom().Hex(),
		})

		// If public transaction
		if txctx.Envelope.GetArgs().GetPrivate() == nil {
			if txctx.Envelope.GetTx().IsSigned() {
				rawTxSender(txctx)
			} else {
				unsignedTxSender(txctx)
			}
		} else {
			protocol := txctx.Envelope.GetProtocol()
			switch {
			case protocol.IsPantheon():
				rawPrivateTxSender(txctx)
			case protocol.IsTessera():
				tesseraRawPrivateTxSender(txctx)
			case protocol.IsConstellation():
				unsignedTxSender(txctx)
			case protocol == nil:
				err := errors.InvalidFormatError(
					"protocol should be specified to send a private transaction",
				).SetComponent(component)
				txctx.Logger.WithError(err).Errorf("sender: could not send private transaction")
				_ = txctx.AbortWithError(err)
			default:
				err := errors.DataError(
					"invalid private protocol %q",
					protocol.String(),
				).SetComponent(component)
				txctx.Logger.WithError(err).Errorf("sender: could not send private transaction")
				_ = txctx.AbortWithError(err)
			}
		}
	}
}
