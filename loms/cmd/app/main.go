package main

import (
	"log"
	"net"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/api/loms_v1"
	"route256/loms/internal/config"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers/cancelorder"
	"route256/loms/internal/handlers/createorder"
	"route256/loms/internal/handlers/listorder"
	"route256/loms/internal/handlers/orderpayed"
	"route256/loms/internal/handlers/stockshandler"
	desc "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	startApp()
}

func startApp() {
	initConfig()
	setupHandles()
	startServer()
	startGRPCServer()
}

func initConfig() {
	//!!!write port numbers with colon ":"
	log.Println("initializing config")

	if err := config.Init(); err != nil {
		log.Fatalln("config init: ", err)
	}
}

func setupHandles() {
	businessLogic := domain.NewService()

	createOrder := createorder.New(businessLogic)
	listOrder := listorder.New(businessLogic)
	orderPayed := orderpayed.New(businessLogic)
	cancelOrder := cancelorder.New(businessLogic)
	stocksHandler := stockshandler.New(businessLogic)

	http.Handle("/createOrder", srvwrapper.New(createOrder.Handle))
	http.Handle("/listOrder", srvwrapper.New(listOrder.Handle))
	http.Handle("/orderPayed", srvwrapper.New(orderPayed.Handle))
	http.Handle("/cancelOrder", srvwrapper.New(cancelOrder.Handle))
	http.Handle("/stocks", srvwrapper.New(stocksHandler.Handle))
}

func startServer() {
	port := config.ConfigData.Port

	log.Println("listening http at: ", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln("cannot listen http: ", err)
	}
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", config.ConfigData.GRPCPort)

	err = errors.WithMessage(err, "grpc server")

	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)
	desc.RegisterLomsServiceServer(s, loms_v1.NewLomsV1(domain.NewService()))

	log.Printf("listening at %v\n", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
