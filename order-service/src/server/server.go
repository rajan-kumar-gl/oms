package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HttpHandler(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// TODO:: Handle Error and send too log manager
		w.Header().Set("Content-Type", "application/json")
		h(w, r, ps)
	}
}
func HttpWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	log.Println(data, statusCode)
	var response = struct {
		StatusCode int         `json:"status"`
		Data       interface{} `json:"data"`
		Err        interface{} `json:"error"`
	}{
		StatusCode: statusCode,
	}

	if statusCode >= 200 && statusCode < 300 {
		response.Data = data
	} else {
		response.Err = data
	}

	responseBytes, err := json.Marshal(response)

	if err != nil {
		log.Printf("error : %+v, res : %+v, %s", err, response, string(responseBytes))
	}
	w.WriteHeader(statusCode)
	w.Write(responseBytes)
}
