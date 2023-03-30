package main

import (
	"context"
	"log"

	"route256/libs/kafka"
	"route256/notifications/config"
	"route256/notifications/internal/reciever"
)

func main() {
	c, err := kafka.NewConsumer(config.ConfigData.Brokers)
	if err != nil {
		log.Fatalln(err)
	}

	topic := config.ConfigData.Topic

	handlers := map[string]reciever.HandleFunc{
		topic: func(id string) {
			log.Println(id)
		},
	}

	r := reciever.NewReciever(c, handlers)
	r.Subscribe(topic)

	<-context.TODO().Done()
}
