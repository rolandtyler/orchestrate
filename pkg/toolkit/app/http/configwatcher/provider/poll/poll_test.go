// +build unit

package poll_test

import (
	"context"
	"testing"
	"time"

	"github.com/consensys/orchestrate/pkg/toolkit/app/http/configwatcher/provider"
	"github.com/consensys/orchestrate/pkg/toolkit/app/http/configwatcher/provider/poll"
	"github.com/consensys/orchestrate/pkg/toolkit/app/http/configwatcher/testutils"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	p := poll.New(
		func(ctx context.Context) (provider.Message, error) {
			return &testutils.Message{Conf: "test-conf"}, nil
		},
		50*time.Millisecond,
	)
	msgs := make(chan provider.Message, 1)
	defer close(msgs)

	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		_ = p.Provide(ctx, msgs)
		close(done)
	}()

	msg := <-msgs
	assert.Equal(t, "test-conf", msg.Configuration(), "Message should have flowed properly")

	msg = <-msgs
	assert.Equal(t, "test-conf", msg.Configuration(), "Message should have flowed properly")

	cancel()
	<-done
}
