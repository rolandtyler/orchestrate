package sarama

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/cenkalti/backoff/v4"
	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
)

type ConsumerDaemon struct {
	client   sarama.Client
	producer sarama.SyncProducer
	group    sarama.ConsumerGroup

	topics  []string
	handler sarama.ConsumerGroupHandler
}

func NewConsumerDaemon(
	client sarama.Client,
	producer sarama.SyncProducer,
	group sarama.ConsumerGroup,
	topics []string,
	handler sarama.ConsumerGroupHandler,
) *ConsumerDaemon {
	return &ConsumerDaemon{
		client:   client,
		producer: producer,
		group:    group,
		topics:   topics,
		handler:  handler,
	}
}

func (d *ConsumerDaemon) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			return backoff.RetryNotify(
				func() error {
					err := d.group.Consume(ctx, d.topics, d.handler)

					// In this case, kafka rebalance was triggered and we want to retry
					if err == nil && ctx.Err() == nil {
						return fmt.Errorf("kafka rebalance was triggered")
					}

					return backoff.Permanent(err)
				},
				backoff.NewConstantBackOff(time.Millisecond*500),
				func(err error, duration time.Duration) {
					log.WithContext(ctx).WithError(err).Warnf("listener: consuming session exited, retrying in %s", duration.String())
				},
			)
		}
	}
}

func (d *ConsumerDaemon) Close() error {
	gr := &multierror.Group{}
	gr.Go(d.producer.Close)
	gr.Go(d.group.Close)
	rerr := gr.Wait()

	err := d.client.Close()
	if err != nil {
		rerr = multierror.Append(rerr, err)
	}

	return rerr.ErrorOrNil()
}
