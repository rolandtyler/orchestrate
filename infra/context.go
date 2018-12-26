package infra

import (
	tracepb "gitlab.com/ConsenSys/client/fr/core-stack/core/protobuf/trace"
	"gitlab.com/ConsenSys/client/fr/core-stack/core/types"
)

// HandlerFunc is base type for a function processing a Trace
type HandlerFunc func(ctx *Context)

// Context allows us to transmit information through middlewares
type Context struct {
	// T stores information about transaction lifecycle in high level types
	T *types.Trace
	// Message that triggered Context execution (typically a sarama.ConsumerMessage)
	Msg interface{}
	// Protobuffer (we attach it to context as an optimization so we can reset it each we re-cycle a context)
	Pb *tracepb.Trace

	// Keys is a key/value pair
	Keys map[string]interface{}

	// Handlers to be executed on context
	handlers []HandlerFunc
	// Handler being executed
	index int
}

// NewContext creates a new context
func NewContext() *Context {
	t := types.NewTrace()
	return &Context{
		Pb:    &tracepb.Trace{},
		T:     t,
		Keys:  make(map[string]interface{}),
		index: -1,
	}
}

// Reset re-initialize context
func (ctx *Context) Reset() {
	ctx.Msg = nil
	ctx.Pb.Reset()
	ctx.T.Reset()
	ctx.Keys = make(map[string]interface{})
	ctx.handlers = nil
	ctx.index = -1
}

// Next should be used in middleware
// It executes pending handlers
func (ctx *Context) Next() {
	ctx.index++
	for s := len(ctx.handlers); ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

// Error attaches an error to context.
func (ctx *Context) Error(err error) *types.Error {
	if err == nil {
		panic("err is nil")
	}

	e, ok := err.(*types.Error)
	if !ok {
		e = &types.Error{
			Err:  err,
			Type: types.ErrorTypeUnknown,
		}
	}
	ctx.T.Errors = append(ctx.T.Errors, e)

	return e
}

// Abort prevents pending handlers to be executed
func (ctx *Context) Abort() {
	ctx.index = len(ctx.handlers)
}

// AbortWithError calls `Abort()` and `Error()``
func (ctx *Context) AbortWithError(err error) *types.Error {
	ctx.Abort()
	return ctx.Error(err)
}

// Prepare re-initializes context, set handlers and set message
func (ctx *Context) Prepare(handlers []HandlerFunc, msg interface{}) {
	ctx.Reset()
	ctx.handlers = handlers
	ctx.Msg = msg
}
