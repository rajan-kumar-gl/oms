package repo

import (
	"errors"

	"github.com/oms/cart-service/src/proto"
)

type implementaion struct {
	// Sample Key value Data store
	container map[int]proto.Cart
}

func NewCart() proto.CartDB {
	//Mocking The DB to in memory storage
	cart := &implementaion{
		container: make(map[int]proto.Cart),
	}
	return cart
}

func (c *implementaion) ByID(userID int) (proto.Cart, error) {
	if c.container == nil {
		return proto.Cart{}, errors.New("unexpected nil storage container found")
	}
	cart, exist := c.container[userID]
	if !exist {
		return proto.Cart{
			Products: map[int]*proto.Product{},
		}, nil
	}
	return cart, nil
}

func (c *implementaion) AddThisToUser(userID int, productID int, howmany int) error {
	if c.container == nil {
		return errors.New("unexpected nil storage container found")
	}

	cart, found := c.container[userID]
	if !found {
		products := map[int]*proto.Product{
			productID: {
				ID:       productID,
				Quantity: howmany,
			},
		}
		c.container[userID] = proto.Cart{Products: products}
		return nil
	}

	_, found = cart.Products[productID]
	if !found {
		cart.Products = map[int]*proto.Product{
			productID: {
				ID:       productID,
				Quantity: howmany,
			},
		}
	} else {
		cart.Products[productID].Quantity += howmany
	}

	c.container[userID] = cart
	return nil

}

func (c *implementaion) IsUserHasProduct(userID, productID int) (bool, error) {
	if c.container == nil {
		return false, errors.New("unexpected nil storage container found")
	}
	cart, found := c.container[userID]
	if !found {
		return false, nil
	}

	if _, found = cart.Products[productID]; found {
		return true, nil
	}

	return false, nil
}
