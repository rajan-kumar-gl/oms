package proto

import "net/http"

type RestAPIClient interface {
	Do(req *http.Request) (resByte []byte, err error)
}
