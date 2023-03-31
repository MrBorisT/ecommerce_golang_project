package main

import (
	"context"
	"fmt"
	"log"
	"route256/libs/kafka"
	"route256/notifications/config"
	"route256/notifications/internal/reciever"
	"route256/notifications/internal/repository"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()
	c, err := kafka.NewConsumer(config.ConfigData.Brokers)
	if err != nil {
		log.Fatalln(err)
	}

	topic := config.ConfigData.Topic

	handlers := map[string]reciever.HandleFunc{
		topic: LogOrderStatusChange,
	}

	pool := OpenDB(ctx)
	defer pool.Close()

	r := reciever.NewReciever(c, repository.NewOffsetsRepo(pool), handlers)
	r.Subscribe(ctx, topic)

	<-ctx.Done()
}

func OpenDB(ctx context.Context) *pgxpool.Pool {
	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.ConfigData.DB.Host,
		config.ConfigData.DB.Port,
		config.ConfigData.DB.User,
		config.ConfigData.DB.Password,
		config.ConfigData.DB.Name,
	)

	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	return pool
}

func LogOrderStatusChange(id string, value []byte) {
	log.Println("order id #", id, "; new status: ", string(value))
}
