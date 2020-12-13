package proto

type ProductUsecase interface {
	GetQty(productID int) (int, error)
	SetQty(productID int, buyingStock int) error
}
