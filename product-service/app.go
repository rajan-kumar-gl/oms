package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/product-service/src/constant"
	"github.com/oms/product-service/src/handler"
	"github.com/oms/product-service/src/repo"
	"github.com/oms/product-service/src/usecase"
)

func main() {

	productRepo := repo.NewProduct()
	log.Print("Successfully INIT Repos")

	productUseCase := usecase.New(productRepo)
	log.Print("Successfully INIT Usecases")

	router := httprouter.New()
	handler.InitProductHandler(router, productUseCase)
	log.Print("Successfully Registered Handller")

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", constant.Server.Host, constant.Server.Port),
	}

	log.Printf("Application server UP : %s:%s", constant.Server.Host, constant.Server.Port)
	log.Printf("Check status on  http://%s:%s/status", constant.Server.Host, constant.Server.Port)
	log.Fatal(server.ListenAndServe())
}
