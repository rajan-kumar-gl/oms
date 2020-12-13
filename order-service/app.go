package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/order-service/src/constant"
	"github.com/oms/order-service/src/handler"
	"github.com/oms/order-service/src/repo"
	"github.com/oms/order-service/src/resource"
	"github.com/oms/order-service/src/usecase"
)

func main() {
	//Initialize resource
	apiResource := resource.New()
	log.Print("Successfully Init API Repos")

	//Initialize repo
	cart := repo.NewCartService(apiResource)
	log.Print("Successfully Init Cart Repos")
	product := repo.NewProductService(apiResource)
	log.Print("Successfully Init Product Repos")
	pl := repo.New()
	log.Print("Successfully Init PL Repos")

	//Initialize use case
	ordermanager := usecase.New(cart, pl, product)
	log.Print("Successfully Init Order Manger Usecases")

	router := httprouter.New()
	handler.InitOMHandler(router, ordermanager)
	log.Print("Successfully RegHandlers")

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", constant.Server.Host, constant.Server.Port),
	}

	log.Printf("Application server UP : %s:%s", constant.Server.Host, constant.Server.Port)
	log.Printf("Check status on  http://%s:%s/status", constant.Server.Host, constant.Server.Port)
	log.Fatal(server.ListenAndServe())
}
