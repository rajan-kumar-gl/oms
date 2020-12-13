package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/product-service/src/proto"
	"github.com/oms/product-service/src/server"
)

func (h handler) getProductQuantity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	productID := ps.ByName("product_id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		server.HttpWriter(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	quantity, err := h.productUseCase.GetQty(id)
	if err != nil {
		server.HttpWriter(w, err, http.StatusBadRequest)
		return
	}

	server.HttpWriter(w, proto.GetProductQuantityResponse{
		ProductID: id,
		AvalQty:   quantity,
	}, http.StatusOK)

	return
}
