package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/product-service/src/proto"
	"github.com/oms/product-service/src/server"
)

type handler struct {
	productUseCase proto.ProductUsecase
}

func InitProductHandler(router *httprouter.Router, productUseCase proto.ProductUsecase) {
	handler := &handler{
		productUseCase: productUseCase,
	}

	router.GET("/status", server.HttpHandler(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"ping\":\"pong\"}"))
	}))

	router.GET("/api/v1/product/quantity/:product_id", server.HttpHandler(handler.getProductQuantity))
	router.POST("/api/v1/product/quantity/:product_id", server.HttpHandler(handler.setProductQuantity))
}
