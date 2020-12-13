package usecases

import "github.com/oms/cart-service/src/proto"

type implemantation struct {
	upstream proto.CartDB
}

func NewCart(repo proto.CartDB) proto.CartUsecase {
	return &implemantation{
		upstream: repo,
	}
}

func (c *implemantation) GetByUserID(userID int) (proto.Cart, error) {
	return c.upstream.ByID(userID)
}

func (c *implemantation) AddProduct(userID int, productID int, buyingStock int) error {
	return c.upstream.AddThisToUser(userID, productID, buyingStock)
}

func (c *implemantation) IsAllReadyInCart(userID, productID int) (bool, error) {
	return c.upstream.IsUserHasProduct(userID, productID)
}
