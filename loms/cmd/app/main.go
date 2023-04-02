package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"route256/libs/kafka"
	"route256/libs/logger"
	"route256/libs/srvwrapper"
	"route256/libs/tracing"
	"route256/libs/transactor"
	"route256/loms/internal/api/loms_v1"
	"route256/loms/internal/config"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers/cancelorder"
	"route256/loms/internal/handlers/createorder"
	"route256/loms/internal/handlers/listorder"
	"route256/loms/internal/handlers/orderpayed"
	"route256/loms/internal/handlers/stockshandler"
	repository "route256/loms/internal/repository/postgres"
	"route256/loms/internal/sender"
	desc "route256/loms/pkg/loms_v1"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	develMode := flag.Bool("devel", false, "developer mode")
	flag.Parse()

	initLogger(*develMode)

	startApp(ctx)

	<-ctx.Done()
}

func startApp(ctx context.Context) {
	initConfig()
	initTracing()
	pool := OpenDB(ctx)
	defer pool.Close()
	service := setupHandlesAndGetService(pool)
	startServer()
	startGRPCServer(service)
}

func initConfig() {
	//!!!write port numbers with colon ":"
	if err := config.Init(); err != nil {
		logger.Fatal("config init", zap.Error(err))
	}
}

func setupHandlesAndGetService(pool *pgxpool.Pool) domain.Service {
	tm := transactor.NewTransactionManager(pool)
	stockRepo := repository.NewStocksRepo(tm)
	orderRepo := repository.NewOrdersRepo(tm)

	producer, err := kafka.NewAsyncProducer(config.ConfigData.Brokers)
	if err != nil {
		logger.Fatal("creating producer", zap.Error(err))
	}

	onSuccess := func(id string) {
		fmt.Println("order success", id)
	}
	onFailed := func(id string) {
		fmt.Println("order failed", id)
	}

	businessLogic := domain.NewService(domain.Deps{
		OrderRepository:    orderRepo,
		StockRepository:    stockRepo,
		TransactionManager: tm,
		StatusSender:       sender.NewStatusSender(producer, config.ConfigData.Topic, onSuccess, onFailed),
	})

	createOrder := createorder.New(businessLogic)
	createOrderHandler := srvwrapper.New(createOrder.Handle)

	listOrder := listorder.New(businessLogic)
	listOrderHandler := srvwrapper.New(listOrder.Handle)

	orderPayed := orderpayed.New(businessLogic)
	orderPayedHandler := srvwrapper.New(orderPayed.Handle)

	cancelOrder := cancelorder.New(businessLogic)
	cancelOrderHandler := srvwrapper.New(cancelOrder.Handle)

	stocks := stockshandler.New(businessLogic)
	stocksHandler := srvwrapper.New(stocks.Handle)

	SetHandler("/createOrder", createOrderHandler)
	SetHandler("/listOrder", listOrderHandler)
	SetHandler("/orderPayed", orderPayedHandler)
	SetHandler("/cancelOrder", cancelOrderHandler)
	SetHandler("/stocks", stocksHandler)

	return businessLogic
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

func startServer() {
	port := config.ConfigData.Port

	logger.Info("server start", zap.String("port", port))

	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}

func startGRPCServer(businessLogic domain.Service) {
	lis, err := net.Listen("tcp", config.ConfigData.GRPCPort)

	err = errors.WithMessage(err, "grpc server")

	if err != nil {
		logger.Fatal("grpc server listen", zap.Error(err))
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
		),
	)

	reflection.Register(s)
	desc.RegisterLomsServiceServer(s, loms_v1.NewLomsV1(businessLogic))

	logger.Info("grpc server started", zap.Any("address", lis.Addr()))

	if err = s.Serve(lis); err != nil {
		logger.Fatal("grpc server error", zap.Error(err))
	}
}

func initLogger(develMode bool) {
	logger.Init(develMode)
}

func SetHandler(route string, handler http.Handler) {
	handler = logger.Middleware(handler)
	handler = tracing.Middleware(handler, route[1:])
	http.Handle(route, handler)
}

func initTracing() {
	tracing.Init("loms service")
}
