package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers/cancelorder"
	"route256/loms/internal/handlers/createorder"
	"route256/loms/internal/handlers/listorder"
	"route256/loms/internal/handlers/orderpayed"
	"route256/loms/internal/handlers/stockshandler"
)

const port = ":8081"

func main() {
	businessLogic := domain.New()
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

	log.Println("listening http at", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
