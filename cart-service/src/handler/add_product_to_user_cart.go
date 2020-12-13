package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/cart-service/src/proto"
	"github.com/oms/cart-service/src/server"
)

func (h handler) addProductToCart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		server.HttpWriter(w, err, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		server.HttpWriter(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(r.FormValue("product_id"))
	if err != nil {
		server.HttpWriter(w, "Invalid ProductID", http.StatusBadRequest)
		return
	}

	buyingStock, err := strconv.Atoi(r.FormValue("qty"))
	if err != nil {
		server.HttpWriter(w, "invalid qty", http.StatusBadRequest)
		return
	}

	err = h.cartUseCase.AddProduct(userID, productID, buyingStock)
	if err != nil {
		server.HttpWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	server.HttpWriter(w, proto.AddProductToCardResponse{
		UserID:    userID,
		ProductID: productID,
		Success:   true,
	}, http.StatusOK)

	return
}
