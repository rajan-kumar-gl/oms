package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/cart-service/src/proto"
	"github.com/oms/cart-service/src/server"
)

type handler struct {
	cartUseCase proto.CartUsecase
}

func InitCartHandler(router *httprouter.Router, cartUseCase proto.CartUsecase) {
	handler := &handler{
		cartUseCase: cartUseCase,
	}

	router.GET("/status", server.HttpHandler(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"ping\":\"pong\"}"))
	}))

	router.GET("/api/v1/cart/status/:user_id/:product_id", server.HttpHandler(handler.isUserAlreadyHasThisProduct))
	router.POST("/api/v1/cart/add/:user_id", server.HttpHandler(handler.addProductToCart))
	router.GET("/api/v1/cart/details/:user_id", server.HttpHandler(handler.getCartInfo))
}
