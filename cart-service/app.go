package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/cart-service/src/constant"
	"github.com/oms/cart-service/src/handler"
	"github.com/oms/cart-service/src/repo"
	"github.com/oms/cart-service/src/usecases"
)

func main() {
	//Initialize repo
	cartRepo := repo.NewCart()
	log.Print("Successfully INIT Repos")
	//Initialize use case
	cartUseCase := usecases.NewCart(cartRepo)
	log.Print("Successfully INIT Usecases")

	router := httprouter.New()
	handler.InitCartHandler(router, cartUseCase)
	log.Print("Successfully Registred Handlers")

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", constant.Server.Host, constant.Server.Port),
	}
	log.Printf("Application server UP : %s:%s", constant.Server.Host, constant.Server.Port)
	log.Printf("Check status on  http://%s:%s/status", constant.Server.Host, constant.Server.Port)

	log.Fatal(server.ListenAndServe())
}
