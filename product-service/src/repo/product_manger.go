package repo

import (
	"errors"
	"fmt"
	"sync"

	"github.com/oms/product-service/src/proto"
)

type implimantation struct {
	inventory map[int]proto.Product
	mutex     *sync.RWMutex
}

//Fix :: Here we are just making dummy product in mem to mock the functionality
func (p *implimantation) addMockProduct() {
	p.inventory = map[int]proto.Product{
		1: {
			ID:       1,
			Name:     "Shirt",
			Amt:      10,
			AvalQty:  50,
			Currency: "INR",
		},
		2: {
			ID:       2,
			Name:     "Book",
			Amt:      10,
			AvalQty:  5,
			Currency: "INR",
		},
		4: {
			ID:       4,
			Name:     "Laptop",
			Amt:      10,
			AvalQty:  60,
			Currency: "INR",
		},
		5: {
			ID:       5,
			Name:     "Mobile",
			Amt:      1000,
			AvalQty:  500,
			Currency: "INR",
		},
		6: {
			ID:       6,
			Name:     "Jacket",
			Amt:      3,
			AvalQty:  50,
			Currency: "INR",
		},
	}

	return
}

func NewProduct() proto.ProductManager {
	product := &implimantation{
		inventory: make(map[int]proto.Product),
		mutex:     &sync.RWMutex{},
	}
	product.addMockProduct()
	return product
}

func (p *implimantation) GetQty(productID int) (int, error) {

	if p.inventory == nil {
		return 0, errors.New("Unexpected empty inventory found")
	}

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if product, found := p.inventory[productID]; found {
		return product.AvalQty, nil
	}
	return 0, errors.New("Invalid Product ID")
}

func (p *implimantation) SetQty(productID int, buyingStock int) error {
	if p.inventory == nil {
		return errors.New("Unexpected empty inventory found")
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	product, found := p.inventory[productID]
	if !found {
		return errors.New("Invalid Product ID")
	}

	remainingQty := product.AvalQty - buyingStock

	if remainingQty < 0 {
		return fmt.Errorf("%d product not avilable", buyingStock)
	}

	product.AvalQty = remainingQty
	p.inventory[productID] = product

	return nil
}
