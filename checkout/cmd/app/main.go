package main

import (
	"context"
	"fmt"
	"log"
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
	"route256/libs/srvwrapper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
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
		log.Fatalln("config init: ", err)
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
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	return pool
}

func setupHandles(lomsConn, productConn *grpc.ClientConn, pool *pgxpool.Pool) {
	lomsClient := loms.NewClient(lomsConn)
	productClient := productsClient.NewClient(productConn, config.ConfigData.Token)
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

	http.Handle("/addToCart", srvwrapper.New(addToCartHandler.Handle))
	http.Handle("/deleteFromCart", srvwrapper.New(deleteFromCart.Handle))
	http.Handle("/listCart", srvwrapper.New(listCart.Handle))
	http.Handle("/purchase", srvwrapper.New(purchase.Handle))
}

func startServer() {
	port := config.ConfigData.Port

	log.Println("listening http at: ", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln("cannot listen http: ", err)
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
		log.Fatalf("cannot connect to loms service: %v\n", err.Error())
	}

	//Product connection
	productConn, err := GetClientConn(config.ConfigData.Services.ProductService)
	if err != nil {
		log.Fatalf("cannot connect to product service: %v\n", err.Error())
	}

	return lomsConn, productConn
}
