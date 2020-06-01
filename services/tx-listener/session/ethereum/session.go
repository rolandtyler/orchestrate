package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/containous/traefik/v2/pkg/log"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/ethclient"
	ethclientutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/ethclient/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/tx"
	evlpstore "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/proto"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-listener/dynamic"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-listener/session"
	hook "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-listener/session/ethereum/hooks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-listener/session/ethereum/offset"
)

const (
	orionPrecompiledContractAddr = "0x000000000000000000000000000000000000007E"
)

type fetchedBlock struct {
	block     *ethtypes.Block
	envelopes []*tx.Envelope
}

type EthClient interface {
	ethclient.ChainLedgerReader
	ethclient.ChainSyncReader
}

type Session struct {
	Chain *dynamic.Chain

	ec    EthClient
	store evlpstore.EnvelopeStoreClient

	hook    hook.Hook
	offsets offset.Manager
	bckOff  backoff.BackOff

	// Listening session
	trigger         chan struct{}
	blockPosition   uint64
	currentChainTip uint64

	// Channel stacking blocks waiting for receipts to be fetched
	fetchedBlocks chan *Future

	errors chan error
}

func (s *Session) Run(ctx context.Context) error {
	err := backoff.RetryNotify(
		func() error {
			err := s.run(ctx)
			if err == context.DeadlineExceeded || err == context.Canceled || ctx.Err() != nil {
				if err == nil {
					err = ctx.Err()
				}

				log.FromContext(ctx).
					WithError(err).
					Info("exiting listener session...")
				return backoff.Permanent(err)
			}

			return err
		},
		s.bckOff,
		func(err error, duration time.Duration) {
			// Print the received error
			log.FromContext(ctx).
				WithError(err).
				Warnf("error in session listener, rebooting in %v...", duration)
		},
	)

	log.FromContext(ctx).Info("listener session exited")
	return err
}

func (s *Session) run(ctx context.Context) (err error) {
	// Initialize session
	err = s.init(ctx)
	if err != nil {
		return err
	}

	cancelableCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start go-routines
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		s.trig()
		s.listen(cancelableCtx)
		wg.Done()
	}()
	go func() {
		s.callHooks(cancelableCtx)
		wg.Done()
	}()

	// Wait for an error or for context to be canceled
	select {
	case err = <-s.errors:
		cancel()
		break
	case <-ctx.Done():
		cancel()
		break
	}

	// We must drain channels before starting a new session
	go func() {
		for e := range s.errors {
			log.FromContext(ctx).
				WithError(e).
				Error("error while listening")
		}
	}()

	log.FromContext(ctx).Debug("waiting for go routines to complete....")
	// Wait for goroutines to complete and close session
	wg.Wait()

	s.close(ctx)
	return err
}

func (s *Session) init(ctx context.Context) error {
	log.FromContext(ctx).Debug("initializing session listener...")

	err := s.initChainID(ctx)
	if err != nil {
		return err
	}

	err = s.initPosition(ctx)
	if err != nil {
		return err
	}

	s.trigger = make(chan struct{}, 1)
	s.errors = make(chan error, 1)
	s.fetchedBlocks = make(chan *Future, 20)

	return nil
}

func (s *Session) initChainID(ctx context.Context) error {
	chain, err := s.ec.Network(ctx, s.Chain.URL)
	if err != nil {
		return err
	}
	s.Chain.ChainID = chain
	return nil
}

func (s *Session) initPosition(ctx context.Context) error {
	blockPosition, err := s.offsets.GetLastBlockNumber(ctx, s.Chain)
	if err != nil {
		return err
	}

	// if blockPosition and startingBlock are different then we have already started listening to that chain
	// and we start at next block since the current one is already processed
	if blockPosition != s.Chain.Listener.StartingBlock {
		blockPosition++
	}

	s.blockPosition = blockPosition

	return nil
}

func (s *Session) listen(ctx context.Context) {
	log.FromContext(ctx).
		WithField("block.start", s.blockPosition).
		Infof("starting fetch block listener")

	ticker := time.NewTicker(s.Chain.Listener.Backoff)
listeningLoop:
	for {
		select {
		case <-ctx.Done():
			log.FromContext(ctx).
				WithField("block.stop", s.blockPosition).
				Debug("stopping fetch block listener")
			break listeningLoop
		case <-s.trigger:
			if (s.currentChainTip > 0) && s.blockPosition <= s.currentChainTip {
				s.fetchedBlocks <- s.fetchBlock(ctx, s.blockPosition)
				s.blockPosition++
				s.trig()
			} else {
				//  We are ahead of chain head so we update chain tip
				tip, err := s.getChainTip(ctx)
				if err != nil {
					s.errors <- err
				} else if tip > s.currentChainTip {
					s.currentChainTip = tip
					s.trig()
				}
			}
		case <-ticker.C:
			s.trig()
		}
	}

	// Close channels
	ticker.Stop()
	close(s.fetchedBlocks)

	log.FromContext(ctx).
		WithField("block.stop", s.blockPosition).
		Infof("fetch block listener has been stopped")
}

func (s *Session) callHooks(ctx context.Context) {
	var err error

	for futureBlock := range s.fetchedBlocks {
		select {
		case res := <-futureBlock.Result():
			// We MUST drain array chan and ignore blocks after an error happened
			if err != nil {
				log.FromContext(ctx).
					WithField("block number", res.(*fetchedBlock).block.NumberU64()).
					Warn("ignoring fetched block")
				continue
			}
			err = s.callHook(ctx, res.(*fetchedBlock))
		case e := <-futureBlock.Err():
			if err == nil && e != nil {
				err = e
			}
		}

		// Close future
		futureBlock.Close()

		if err != nil {
			s.errors <- err
		}
	}

	log.FromContext(ctx).Debug("call hooks loop has been stopped")
}

func (s *Session) callHook(ctx context.Context, block *fetchedBlock) error {
	// Call hook
	err := s.hook.AfterNewBlock(ctx, s.Chain, block.block, block.envelopes)
	if err == nil {
		// Update last block processed
		err = s.offsets.SetLastBlockNumber(ctx, s.Chain, block.block.NumberU64())
	}
	return err
}

func (s *Session) fetchBlock(ctx context.Context, blockPosition uint64) *Future {
	return NewFuture(func() (interface{}, error) {
		blck, err := s.ec.BlockByNumber(
			ethclientutils.RetryNotFoundError(ctx, true),
			s.Chain.URL,
			big.NewInt(int64(blockPosition)),
		)

		if err != nil {
			log.FromContext(ctx).
				WithError(err).
				WithField("block.number", blockPosition).
				Errorf("failed to fetch block")

			return nil, errors.ConnectionError(err.Error())
		}

		block := &fetchedBlock{block: blck}

		ctx = multitenancy.WithTenantID(ctx, s.Chain.TenantID)

		// Fetch envelopes from the envelope store
		envelopeMap, err := s.fetchEnvelopes(ctx, blck.Transactions())
		if err != nil {
			return nil, err
		}

		// Fetch receipt for transactions
		futureEnvelopes := s.fetchReceipts(ctx, blck.Transactions(), envelopeMap)

		// Await receipts
		block.envelopes, err = awaitReceipts(futureEnvelopes)
		if err != nil {
			return nil, err
		}

		return block, nil
	})
}

func (s *Session) fetchEnvelopes(ctx context.Context, transactions ethtypes.Transactions) (envelopeMap map[string]*tx.Envelope, err error) {
	envelopeMap = make(map[string]*tx.Envelope)

	// Load envelopes from the envelope store
	if len(transactions) > 0 {
		var txHashes []string
		for _, t := range transactions {
			txHashes = append(txHashes, t.Hash().String())
		}
		var resp *evlpstore.LoadByTxHashesResponse
		resp, err = s.store.LoadByTxHashes(
			ctx,
			&evlpstore.LoadByTxHashesRequest{
				ChainId:  s.Chain.ChainID.String(),
				TxHashes: txHashes,
			})
		if err != nil {
			return nil, err
		}
		for _, t := range resp.Responses {
			envelope, er := t.GetEnvelope().Envelope()
			if er != nil {
				return nil, er
			}

			// Filter by the envelopes belonging to same session CHAIN_UUID
			if envelope.ChainUUID == s.Chain.UUID {
				envelopeMap[t.Envelope.GetTxHash()] = envelope
			}
		}
	}

	return envelopeMap, nil
}

func (s *Session) fetchReceipts(ctx context.Context, transaction ethtypes.Transactions, envelopeMap map[string]*tx.Envelope) (futureEnvelopes []*Future) {
	for _, blckTx := range transaction {
		switch {
		case isPrivTx(blckTx) && isInternalTx(envelopeMap, blckTx):
			futureEnvelopes = append(futureEnvelopes, s.fetchPrivateReceipt(
				ctx,
				envelopeMap[blckTx.Hash().String()],
				blckTx.Hash()))
			continue
		case isInternalTx(envelopeMap, blckTx):
			futureEnvelopes = append(futureEnvelopes, s.fetchReceipt(
				ctx,
				envelopeMap[blckTx.Hash().String()],
				blckTx.Hash()))
			continue
		case isPrivTx(blckTx) && s.Chain.Listener.ExternalTxEnabled:
			futureEnvelopes = append(futureEnvelopes, s.fetchPrivateReceipt(
				ctx,
				tx.NewEnvelope().SetTxHash(blckTx.Hash()).SetChainName(s.Chain.Name),
				blckTx.Hash()))
			continue
		case s.Chain.Listener.ExternalTxEnabled:
			futureEnvelopes = append(futureEnvelopes, s.fetchReceipt(
				ctx,
				tx.NewEnvelope().SetTxHash(blckTx.Hash()).SetChainName(s.Chain.Name),
				blckTx.Hash()))
			continue
		default:
			continue
		}
	}
	return futureEnvelopes
}

func awaitReceipts(futureEnvelopes []*Future) (envelopes []*tx.Envelope, err error) {
	for _, futureEnvelope := range futureEnvelopes {
		select {
		case e := <-futureEnvelope.Err():
			if err == nil {
				err = e
			}
		case res := <-futureEnvelope.Result():
			envelopes = append(envelopes, res.(*tx.Envelope))
		}

		// Close future
		futureEnvelope.Close()
	}
	if err != nil {
		return nil, err
	}
	return envelopes, nil
}

func isInternalTx(envelopeMap map[string]*tx.Envelope, transaction *ethtypes.Transaction) bool {
	_, ok := envelopeMap[transaction.Hash().String()]
	return ok
}

func isPrivTx(transaction *ethtypes.Transaction) bool {
	// A enclavekey tx has as To address the pre-deployed smart-contract
	return transaction.To() != nil && transaction.To().String() == orionPrecompiledContractAddr
}

func (s *Session) fetchReceipt(ctx context.Context, envelope *tx.Envelope, txHash ethcommon.Hash) *Future {
	return NewFuture(func() (interface{}, error) {
		log.FromContext(ctx).
			WithField("tx.hash", txHash.Hex()).
			WithField("chainUUID", s.Chain.UUID).
			Debug("fetching fetch receipt")

		receipt, err := s.ec.TransactionReceipt(
			ethclientutils.RetryNotFoundError(ctx, true),
			s.Chain.URL,
			txHash,
		)

		if err != nil {
			log.FromContext(ctx).
				WithError(err).
				WithField("tx.hash", txHash.Hex()).
				WithField("chainUUID", s.Chain.UUID).
				Errorf("failed to fetch receipt")

			return nil, err
		}

		// Attach receipt to envelope
		return envelope.SetReceipt(receipt.
			SetBlockHash(ethcommon.HexToHash(receipt.GetBlockHash())).
			SetBlockNumber(receipt.GetBlockNumber()).
			SetTxIndex(receipt.TxIndex)), nil
	})
}

func (s *Session) fetchPrivateReceipt(ctx context.Context, envelope *tx.Envelope, txHash ethcommon.Hash) *Future {
	return NewFuture(func() (interface{}, error) {
		if envelope == nil {
			err := fmt.Errorf("envelope cannot be nil")
			log.FromContext(ctx).
				WithError(err).
				Errorf("envelope cannot be nil")
			return nil, err
		}

		log.FromContext(ctx).
			WithField("tx.hash", txHash.Hex()).
			WithField("chainUUID", s.Chain.UUID).
			Debug("fetching private receipt")

		receipt, err := s.ec.PrivateTransactionReceipt(
			ethclientutils.RetryNotFoundError(ctx, true),
			s.Chain.URL,
			txHash,
		)

		// We exit ONLY when we even failed to fetch the marking tx receipt, otherwise
		// error is being appended to the envelope
		if err != nil && receipt == nil {
			log.FromContext(ctx).
				WithError(err).
				WithField("tx.hash", txHash.Hex()).
				WithField("chainUUID", s.Chain.UUID).
				Error("failed to fetch receipt")

			return nil, err
		} else if err != nil {
			log.FromContext(ctx).
				WithError(err).
				WithField("tx.hash", txHash.Hex()).
				WithField("chainUUID", s.Chain.UUID).
				Warn("failed to fetch private receipt")
		}

		log.FromContext(ctx).
			WithField("TxHash", receipt.TxHash).
			WithField("PrivateFrom", receipt.PrivateFrom).
			WithField("PrivateFor", receipt.PrivateFor).
			WithField("PrivacyGroupID", receipt.PrivacyGroupId).
			WithField("Status", receipt.Status).
			Debug("private Receipt fetched")

		// Bind the hybrid receipt to the envelope
		return envelope.SetReceipt(receipt.
			SetBlockHash(ethcommon.HexToHash(receipt.GetBlockHash())).
			SetBlockNumber(receipt.GetBlockNumber()).
			SetTxHash(txHash).
			SetTxIndex(receipt.TxIndex)), nil
	})
}

func (s *Session) getChainTip(ctx context.Context) (tip uint64, err error) {
	head, err := s.ec.HeaderByNumber(
		ethclientutils.RetryNotFoundError(ctx, true),
		s.Chain.URL,
		nil,
	)
	if err != nil {
		log.FromContext(ctx).WithError(err).Errorf("failed to fetch chain head")
		return 0, err
	}

	if head.Number.Uint64() > s.Chain.Listener.Depth {
		tip = head.Number.Uint64() - s.Chain.Listener.Depth
	}

	return
}

func (s *Session) trig() {
	select {
	case s.trigger <- struct{}{}:
	default:
		// already triggered
	}
}

func (s *Session) close(ctx context.Context) {
	log.FromContext(ctx).Debug("closing listener session...")
	close(s.errors)
	close(s.trigger)
}

type SessionBuilder struct {
	hook    hook.Hook
	offsets offset.Manager

	ec    EthClient
	store evlpstore.EnvelopeStoreClient
}

func NewSessionBuilder(hk hook.Hook, offsets offset.Manager, ec EthClient, store evlpstore.EnvelopeStoreClient) *SessionBuilder {
	return &SessionBuilder{
		hook:    hk,
		offsets: offsets,
		ec:      ec,
		store:   store,
	}
}

func (b *SessionBuilder) NewSession(chain *dynamic.Chain) (session.Session, error) {
	return b.newSession(chain), nil
}

func (b *SessionBuilder) newSession(chain *dynamic.Chain) *Session {
	return &Session{
		Chain:   chain,
		ec:      b.ec,
		store:   b.store,
		hook:    b.hook,
		offsets: b.offsets,
		bckOff:  backoff.NewConstantBackOff(2 * time.Second),
	}
}
