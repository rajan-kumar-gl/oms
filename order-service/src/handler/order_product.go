package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/order-service/src/proto"
	"github.com/oms/order-service/src/server"
)

func (h handler) orderProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		server.HttpWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(r.Header.Get("HA-UserId"))
	if err != nil {
		server.HttpWriter(w, "invalid user id", http.StatusBadRequest)
		return
	}

	productId, err := strconv.Atoi(ps.ByName("product_id"))
	if err != nil {
		server.HttpWriter(w, "invalid product id", http.StatusBadRequest)
		return
	}

	howMany, err := strconv.Atoi(r.FormValue("qty"))
	if err != nil {
		server.HttpWriter(w, "invalid qty stock", http.StatusBadRequest)
		return
	}

	err = h.orderUsecase.OrderProduct(userId, productId, howMany)
	if err != nil {
		server.HttpWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	server.HttpWriter(w, proto.OrderInfo{
		UserID:      userId,
		ProductID:   productId,
		BuyingStock: howMany,
		Success:     true,
	}, http.StatusOK)
	return
}
