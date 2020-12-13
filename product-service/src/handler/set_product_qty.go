package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oms/product-service/src/proto"
	"github.com/oms/product-service/src/server"
)

func (h handler) setProductQuantity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if err := r.ParseForm(); err != nil {
		server.HttpWriter(w, err, http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(ps.ByName("product_id"))
	if err != nil {
		server.HttpWriter(w, "invalid product id", http.StatusInternalServerError)
		return
	}

	qty, err := strconv.Atoi(r.FormValue("qty"))
	if err != nil {
		server.HttpWriter(w, "invalid qty", http.StatusInternalServerError)
		return
	}

	err = h.productUseCase.SetQty(id, qty)
	if err != nil {
		server.HttpWriter(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	server.HttpWriter(w, proto.SetProductQuantityResponse{
		ProductID: id,
		Success:   true,
	}, http.StatusOK)

	return
}
