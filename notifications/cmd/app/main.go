package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"route256/libs/kafka"
	"route256/libs/logger"
	"route256/notifications/config"
	"route256/notifications/internal/reciever"
	"route256/notifications/internal/repository"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	develMode := flag.Bool("devel", false, "developer mode")
	flag.Parse()

	initLogger(*develMode)

	c, err := kafka.NewConsumer(config.ConfigData.Brokers)
	if err != nil {
		logger.Fatal("creating consumer", zap.Error(err))
	}

	topic := config.ConfigData.Topic

	handlers := map[string]reciever.HandleFunc{
		topic: LogOrderStatusChange,
	}

	pool := OpenDB(ctx)
	defer pool.Close()

	r := reciever.NewReciever(c, repository.NewOffsetsRepo(pool), handlers)
	if err = r.Subscribe(ctx, topic); err != nil {
		log.Fatal("subscription failed", zap.Error(err))
	}

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
		logger.Fatal("db connect", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Fatal("db ping", zap.Error(err))
	}

	return pool
}

func LogOrderStatusChange(id string, value []byte) {
	logger.Info("order change", zap.String("id", id), zap.String("status", string(value)))
}

func initLogger(develMode bool) {
	logger.Init(develMode)
}
