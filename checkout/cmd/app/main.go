package main

import (
	"log"
	"net/http"
	"route256/checkout/internal/clients/loms"
	product "route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers/addtocart"
	deletefromcart "route256/checkout/internal/handlers/deleteFromCart"
	listcart "route256/checkout/internal/handlers/listCart"
	"route256/checkout/internal/handlers/purchase"
	"route256/libs/srvwrapper"
)

func main() {
	initConfig()
	setupHandles()
	startServer()
}

func initConfig() {
	err := config.Init()
	if err != nil {
		log.Fatalln("config init: ", err)
	}
}

func setupHandles() {
	lomsClient := loms.New(config.ConfigData.Services.Loms)
	productClient := product.New(config.ConfigData.Services.ProductService, config.ConfigData.Token)

	businessLogic := domain.New(lomsClient, productClient)

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
