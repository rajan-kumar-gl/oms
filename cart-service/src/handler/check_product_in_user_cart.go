package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/cart-service/src/proto"
	"github.com/oms/cart-service/src/server"
)

func (h handler) isUserAlreadyHasThisProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		server.HttpWriter(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		server.HttpWriter(w, "invalid user id", http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(ps.ByName("product_id"))
	if err != nil {
		server.HttpWriter(w, "invalid product id", http.StatusBadRequest)
		return
	}

	yes, err := h.cartUseCase.IsAllReadyInCart(userID, productID)
	if err != nil {
		server.HttpWriter(w, "error checking cart details", http.StatusInternalServerError)
		return
	}

	response := proto.UserProduct{
		UserID:     userID,
		ProductID:  productID,
		CartStatus: yes,
	}
	server.HttpWriter(w, response, http.StatusOK)
	return
}
