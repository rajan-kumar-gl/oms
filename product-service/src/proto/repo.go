package proto

type ProductManager interface {
	GetQty(productID int) (int, error)
	SetQty(productID int, buyingStock int) error
}
