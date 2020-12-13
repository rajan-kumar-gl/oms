package usecase

import "github.com/oms/product-service/src/proto"

type implemantation struct {
	repo proto.ProductManager
}

func New(repo proto.ProductManager) proto.ProductManager {
	return &implemantation{
		repo: repo,
	}
}

func (p *implemantation) GetQty(productID int) (int, error) {
	return p.repo.GetQty(productID)
}

func (p *implemantation) SetQty(productID int, buyingStock int) error {
	return p.repo.SetQty(productID, buyingStock)
}
