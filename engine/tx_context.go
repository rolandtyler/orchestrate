package engine

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/envelope"
	err "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/error"
)

// TxContext is the most important part of an engine.
// It allows to pass variables between handlers
type TxContext struct {
	// Envelope stores all information about transaction lifecycle
	Envelope *envelope.Envelope

	// Message that triggered TxContext execution (typically a sarama.ConsumerMessage)
	Msg Msg

	// Array of sequences of handlers to execute on a given context
	stack []*sequence

	// Logger logrus log entry for this TxContext execution
	Logger *log.Entry

	// ctx is a go context that is attached to the TxContext
	// It allows to carry deadlines, cancelation signals, ettxctx. between handlers
	//
	// This approach is not recommended by go context documentation
	// txctx.f. https://golang.org/pkg/context/#pkg-overview
	//
	// Still this recommendation against has been actively questioned
	// (txctx.f https://github.com/golang/go/issues/22602)
	// Also net/http has been following this implementation for the Request object
	// (txctx.f. https://github.com/golang/go/blob/master/src/net/http/request.go#L107)
	ctx context.Context
}

// NewTxContext creates a new TxContext
func NewTxContext() *TxContext {
	return &TxContext{
		Envelope: &envelope.Envelope{},
	}
}

// Reset re-initialize TxContext
func (txctx *TxContext) Reset() {
	txctx.ctx = nil
	txctx.Msg = nil
	txctx.Envelope.Reset()
	txctx.stack = nil
	txctx.Logger = nil
}

// Next should be used only inside middleware
// It executes the pending handlers in the chain inside the calling handler
func (txctx *TxContext) Next() {
	if len(txctx.stack) > 0 {
		txctx.stack[len(txctx.stack)-1].next()
	}
}

// Error attaches an error to TxContext
func (txctx *TxContext) Error(e error) *err.Error {
	if e == nil {
		panic("err is nil")
	}

	rv := err.FromError(e)
	txctx.Envelope.Errors = append(txctx.Envelope.Errors, rv)

	return rv
}

// Abort prevents pending handlers to be executed
func (txctx *TxContext) Abort() {
	for s := len(txctx.stack) - 1; s >= 0; s-- {
		txctx.stack[s].abort()
	}
}

// AbortWithError calls `Abort()` and `Error()``
func (txctx *TxContext) AbortWithError(e error) *err.Error {
	txctx.Abort()
	return txctx.Error(e)
}

// Prepare re-initializes TxContext, set handlers, set logger and set message
func (txctx *TxContext) Prepare(logger *log.Entry, msg Msg) *TxContext {
	txctx.Reset()
	txctx.Msg = msg
	txctx.Logger = logger
	return txctx
}

type txCtxKey string

// Set is used to store a new key/value pair exclusively for this context
func (txctx *TxContext) Set(key string, value interface{}) {
	txctx.WithContext(context.WithValue(txctx.Context(), txCtxKey(key), value))
}

// Get returns the value for the given key
func (txctx *TxContext) Get(key string) interface{} {
	return txctx.Context().Value(txCtxKey(key))
}

// Context returns the go context attached to TxContext.
// To change the go context, use WithContext.
//
// The returned context is always non-nil; it defaults to the background context.
func (txctx *TxContext) Context() context.Context {
	if txctx.ctx != nil {
		return txctx.ctx
	}
	return context.Background()
}

// WithContext attach a go context to TxContext
// The go context provided as argument must be non nil or WithContext will panic
func (txctx *TxContext) WithContext(ctx context.Context) *TxContext {
	if ctx == nil {
		panic("nil context")
	}
	txctx.ctx = ctx
	return txctx
}

func (txctx *TxContext) applyHandlers(handlers ...HandlerFunc) {
	// Recycle sequence
	seq := seqPool.Get().(*sequence)
	defer seqPool.Put(seq)

	// Initialize sequence
	seq.index = -1
	seq.handlers = handlers
	seq.txctx = txctx

	// Attach the sequence to the TxContext
	txctx.stack = append(txctx.stack, seq)

	// Execute sequence
	seq.next()

	// Once executed remove the sequence
	txctx.stack = txctx.stack[:len(txctx.stack)-1]
}

type sequence struct {
	// chain of handlers to be executed in the sequence
	handlers []HandlerFunc

	// index of the handler being executed
	index int

	// context the sequence is attached to
	txctx *TxContext
}

// sequences are pooled to relieve presure on garbage collector
var seqPool = sync.Pool{
	New: func() interface{} { return &sequence{index: -1} },
}

func (seq *sequence) next() {
	seq.index++
	for s := len(seq.handlers); seq.index < s; seq.index++ {
		seq.handlers[seq.index](seq.txctx)
	}
}

func (seq *sequence) abort() {
	seq.index = len(seq.handlers)
}

func CombineHandlers(handlers ...HandlerFunc) HandlerFunc {
	return func(txctx *TxContext) {
		txctx.applyHandlers(handlers...)
	}
}

// Msg is an abstract interface supported by any kind of message handled by the engine
type Msg interface {
	// Entrypoint returns an indication on where the message comes from
	Entrypoint() string
}
