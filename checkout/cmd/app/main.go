package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"route256/checkout/internal/clients/grpc/loms"
	productsClient "route256/checkout/internal/clients/grpc/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers/addtocart"
	deletefromcart "route256/checkout/internal/handlers/deleteFromCart"
	listcart "route256/checkout/internal/handlers/listCart"
	"route256/checkout/internal/handlers/purchase"
	repository "route256/checkout/internal/repository/postgres"
	productServiceAPI "route256/checkout/pkg/product"
	"route256/libs/logger"
	"route256/libs/srvwrapper"
	"route256/libs/tracing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const DEFAULT_RPS = 10

func main() {
	develMode := flag.Bool("devel", false, "developer mode")
	flag.Parse()

	initLogger(*develMode)
	initTracing()
	initConfig()

	lomsConn, productConn := ConnectToGRPCServices()
	defer CloseConnections(lomsConn, productConn)

	pool := OpenDB()
	defer pool.Close()

	setupHandles(lomsConn, productConn, pool)
	startServer()
}

func initConfig() {
	err := config.Init()
	if err != nil {
		logger.Fatal("config init", zap.Error(err))
	}
}

func OpenDB() *pgxpool.Pool {
	ctx := context.Background()

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

func setupHandles(lomsConn, productConn *grpc.ClientConn, pool *pgxpool.Pool) {
	lomsClient := loms.NewClient(lomsConn)
	psClient := productServiceAPI.NewProductServiceClient(productConn)
	limiter := rate.NewLimiter(rate.Every(time.Second*1), DEFAULT_RPS)
	deps := productsClient.Deps{
		ProductClient: psClient,
		Token:         config.ConfigData.Token,
		Limiter:       limiter,
	}
	productClient := productsClient.NewClient(deps)
	repository := repository.NewCartRepo(pool)

	businessLogic := domain.NewCheckoutService(domain.Deps{
		LOMS:           lomsClient,
		ProductChecker: productClient,
		CartRepository: repository,
	})

	addToCartHandler := addtocart.New(businessLogic)
	deleteFromCart := deletefromcart.New(businessLogic)
	listCart := listcart.New(businessLogic)
	purchase := purchase.New(businessLogic)

	SetHandler("/addToCart", srvwrapper.New(addToCartHandler.Handle))
	SetHandler("/deleteFromCart", srvwrapper.New(deleteFromCart.Handle))
	SetHandler("/listCart", srvwrapper.New(listCart.Handle))
	SetHandler("/purchase", srvwrapper.New(purchase.Handle))
}

func startServer() {
	port := config.ConfigData.Port

	logger.Info("start server", zap.String("port", port))

	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}

func GetClientConn(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.WithMessage(err, "grpc dial")
	}

	return conn, nil
}

func CloseConnections(connections ...*grpc.ClientConn) {
	for _, connection := range connections {
		connection.Close()
	}
}

func ConnectToGRPCServices() (*grpc.ClientConn, *grpc.ClientConn) {
	//LOMS connection
	lomsConn, err := GetClientConn(config.ConfigData.Services.Loms)
	if err != nil {
		logger.Fatal("loms connect", zap.Error(err))
	}

	//Product connection
	productConn, err := GetClientConn(config.ConfigData.Services.ProductService)
	if err != nil {
		logger.Fatal("product service", zap.Error(err))
	}

	return lomsConn, productConn
}

func initLogger(develMode bool) {
	logger.Init(develMode)
}

func initTracing() {
	tracing.Init("checkout")
}

func SetHandler(route string, handler http.Handler) {
	handler = logger.Middleware(handler)
	handler = tracing.Middleware(handler, route[1:])
	http.Handle(route, handler)
}
