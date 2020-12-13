package proto

type CartUsecase interface {
	IsAllReadyInCart(userID, productID int) (bool, error)
	GetByUserID(userID int) (Cart, error)
	AddProduct(userID int, productID int, buyingStock int) error
}
