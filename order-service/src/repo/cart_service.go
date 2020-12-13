package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/oms/order-service/src/constant"
	"github.com/oms/order-service/src/proto"
)

type implemantation struct {
	client proto.RestAPIClient
}

func NewCartService(client proto.RestAPIClient) proto.CartService {
	return &implemantation{
		client: client,
	}
}

func (r *implemantation) IsProductInCart(userID int, productID int) (bool, error) {
	var (
		userProductResponse = proto.CartServiceResponse{}
	)

	uri := fmt.Sprintf("%v/status/%v/%v", constant.CartServiceEnpoint, userID, productID)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return false, fmt.Errorf("unable to prepare request for %s", uri)
	}

	req.Header.Set("Content-Type", "application/json")
	respBytes, err := r.client.Do(req)
	if err != nil {
		log.Printf("CartService::CheckProductInCart Error : %+v", err)
		return false, errors.New("getting error from upstream : cart service")
	}

	err = json.Unmarshal(respBytes, &userProductResponse)
	if err != nil {
		log.Printf("CartService::CheckProductInCart Unmarshal Error : %+v", err)
		return false, errors.New("CartService.unable to parse upstream error")
	}
	return userProductResponse.Data.CartStatus, nil
}
