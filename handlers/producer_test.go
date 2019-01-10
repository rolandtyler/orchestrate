package handlers

import (
	"fmt"
	"math/big"
	"sync"
	"testing"

	tracepb "gitlab.com/ConsenSys/client/fr/core-stack/core/protobuf/trace"
	"gitlab.com/ConsenSys/client/fr/core-stack/core/types"
)

type MockTraceProducer struct {
	t *testing.T
}

func (p *MockTraceProducer) Produce(pb *tracepb.Trace) error {
	if pb.GetChain().GetId() == "0x0" {
		return fmt.Errorf("Could not produce")
	}
	return nil
}

func makeProducerContext(i int) *types.Context {
	ctx := types.NewContext()
	ctx.Reset()
	switch i % 2 {
	case 0:
		ctx.T.Chain().ID = big.NewInt(0)
		ctx.Keys["errors"] = 1
	case 1:
		ctx.T.Chain().ID = big.NewInt(10)
		ctx.Keys["errors"] = 0
	}
	return ctx
}

func TestProducer(t *testing.T) {
	mp := MockTraceProducer{t: t}
	producer := Producer(&mp)

	rounds := 100
	outs := make(chan *types.Context, rounds)
	wg := &sync.WaitGroup{}
	for i := 0; i < rounds; i++ {
		wg.Add(1)
		ctx := makeProducerContext(i)
		go func(ctx *types.Context) {
			defer wg.Done()
			producer(ctx)
			outs <- ctx
		}(ctx)
	}
	wg.Wait()
	close(outs)

	if len(outs) != rounds {
		t.Errorf("Marker: expected %v outs but got %v", rounds, len(outs))
	}

	for out := range outs {
		errCount := out.Keys["errors"].(int)
		if len(out.T.Errors) != errCount {
			t.Errorf("Marker: expected %v errors but got %v", errCount, out.T.Errors)
		}
	}
}
