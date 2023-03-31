package sender

import (
	"encoding/json"
	"fmt"
	"time"

	"log"

	"github.com/Shopify/sarama"
)

type sender struct {
	producer sarama.AsyncProducer
	topic    string
}

type OrderStatus struct {
	Status string `json:"status"`
}

type Handler func(id string)

func NewStatusSender(producer sarama.AsyncProducer, topic string, onSuccess, onFailed Handler) *sender {
	s := &sender{
		producer: producer,
		topic:    topic,
	}

	// config.Producer.Return.Errors = true
	go func() {
		for e := range producer.Errors() {
			bytes, _ := e.Msg.Key.Encode()

			onFailed(string(bytes))
			fmt.Println(e.Msg.Key, e.Error())
		}
	}()

	// config.Producer.Return.Successes = true
	go func() {
		for m := range producer.Successes() {
			bytes, _ := m.Key.Encode()
			onSuccess(string(bytes))
			log.Printf("order id: %s, partition: %d, offset: %d", string(bytes), m.Partition, m.Offset)
		}
	}()

	return s
}

func (s *sender) SendStatusChange(orderID int64, status string) {
	bytes, err := json.Marshal(OrderStatus{
		Status: status,
	})
	if err != nil {
		log.Println("sending status change", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(bytes),
		Key:       sarama.StringEncoder(fmt.Sprint(orderID)),
		Timestamp: time.Now(),
	}

	s.producer.Input() <- msg
}
