package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/cart-service/src/proto"
	"github.com/oms/cart-service/src/server"
)

func (h handler) getCartInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		server.HttpWriter(w, "Invalid User ID", http.StatusInternalServerError)
		return
	}

	cart, err := h.cartUseCase.GetByUserID(id)
	if err != nil {
		server.HttpWriter(w, err, http.StatusInternalServerError)
		return
	}

	server.HttpWriter(w, proto.CartInfoResponse{
		UserID: id,
		Cart:   cart,
	}, http.StatusOK)

	return
}
