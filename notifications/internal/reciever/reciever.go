package reciever

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type HandleFunc func(id string, value []byte)

type OffsetRepository interface {
	CreateRepo(ctx context.Context, partitionID int32) error
	GetOffsetForRepo(ctx context.Context, partitionID int32) (int64, error)
	UpdateOffset(ctx context.Context, partitionID int32, newOffset int64) error
	GetOffsets(ctx context.Context) (map[int32]int64, error)
}

type Reciever struct {
	consumer   sarama.Consumer
	offsetRepo OffsetRepository
	handlers   map[string]HandleFunc
}

func NewReciever(consumer sarama.Consumer, offsetRepo OffsetRepository, handlers map[string]HandleFunc) *Reciever {
	return &Reciever{
		consumer:   consumer,
		handlers:   handlers,
		offsetRepo: offsetRepo,
	}
}

func (r *Reciever) Subscribe(ctx context.Context, topic string) error {
	handler := r.handlers[topic]

	partitionList, err := r.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	offsets, err := r.offsetRepo.GetOffsets(ctx)
	if err != nil {
		return errors.WithMessage(err, "getting initial offsets")
	}

	for _, partition := range partitionList {
		initialOffset, ok := offsets[partition]
		if !ok {
			r.offsetRepo.CreateRepo(ctx, partition)
		}

		pc, err := r.consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			for message := range pc.Messages() {
				k := string(message.Key)
				handler(k, message.Value)
				if err := r.offsetRepo.UpdateOffset(ctx, partition, message.Offset); err != nil {
					log.Println("error updating #", partition, " partition, offset #", message.Offset, ": ", err)
				}
			}
		}(pc, partition)
	}

	return nil
}
