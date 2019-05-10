package main

import (
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

var kafkaUrls = []string{"localhost:9092"}
var inTopic = "topic-tx-nonce"

func main() {
	consumer, err := sarama.NewConsumer(kafkaUrls, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = consumer.Close(); err != nil {
			log.WithError(err).Fatal("e2e: could not close consumer")
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(inTopic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.WithFields(log.Fields{
				"offset": msg.Offset,
			}).Infof("Consumed message offset")
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
