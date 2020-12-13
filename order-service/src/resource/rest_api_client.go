package resource

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/oms/order-service/src/proto"
)

type implemantation struct {
	httpClient *http.Client
}

func New() proto.RestAPIClient {
	return &implemantation{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (r *implemantation) Do(req *http.Request) (resByte []byte, err error) {
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
