package proto

type OrderManager interface {
	OrderProduct(userID int, productID int, stock int) error
}
