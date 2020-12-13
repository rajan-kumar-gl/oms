package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/order-service/src/proto"
	"github.com/oms/order-service/src/server"
)

type handler struct {
	orderUsecase proto.OrderManager
}

func InitOMHandler(router *httprouter.Router, orderusecase proto.OrderManager) {
	handler := &handler{
		orderUsecase: orderusecase,
	}

	router.GET("/status", server.HttpHandler(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"ping\":\"pong\"}"))
	}))
	router.POST("/api/angel-broking/om/:product_id", server.HttpHandler(handler.orderProduct))
}
