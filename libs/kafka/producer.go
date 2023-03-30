package kafka

import (
	"github.com/Shopify/sarama"
)

func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner // sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Idempotent = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func NewAsyncProducer(brokers []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner // sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Idempotent = true // exactly once

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
