package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/oms/order-service/src/constant"
	"github.com/oms/order-service/src/proto"
)

func NewProductService(client proto.RestAPIClient) proto.ProductService {
	return &implemantation{
		client: client,
	}
}

func (r *implemantation) GetQtyOfProduct(productID int) (int, error) {
	var (
		quantity                int
		productQuantityResponse = proto.GetQuantityResponse{}
	)

	url := fmt.Sprintf("%v/quantity/%v", constant.ProductServiceEnpoint, productID)
	fmt.Printf(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("ProductService.GetProductQuantity error while creating req: ", err)
		return quantity, errors.New("GetProductQuantity.unable to prepare request")
	}

	req.Header.Set("Content-Type", "application/json")
	respBytes, err := r.client.Do(req)
	if err != nil {
		log.Println("ProductService.GetProductQuantity error in response: ", err)
		return quantity, errors.New("error getting product quantity")
	}

	err = json.Unmarshal(respBytes, &productQuantityResponse)
	if err != nil {
		log.Println("ProductService.GetProductQuantity  error while unmarshalling upstream response: ", err)
		return quantity, errors.New("unable to proccesd upstream error")
	}
	return productQuantityResponse.Data.Qty, nil
}

func (r *implemantation) LockProduct(productID int, howmany int) error {
	var (
		productQuantityResponse = proto.SetQuantityResponse{}
	)
	uri := fmt.Sprintf("%v/quantity/%v", constant.ProductServiceEnpoint, productID)
	data := url.Values{}
	data.Set("qty", fmt.Sprintf("%v", howmany))

	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println("ProductService.SetProductQuantity error while creating request : ", err)
		return errors.New("error creating set quantity request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	respBytes, err := r.client.Do(req)
	if err != nil {
		log.Println("ProductService.SetProductQuantity error in upstream  : ", err)
		return errors.New("error in  setting product quantity")
	}

	err = json.Unmarshal(respBytes, &productQuantityResponse)
	if err != nil {
		log.Println("ProductService. SetProductQuantity error while unmarshalling : ", err)
		return errors.New("error in setting  product quantity")
	}
	if !productQuantityResponse.Data.Success {
		return errors.New("unable to update qty.")
	}
	return nil
}
