package proto

type CartService interface {
	IsProductInCart(userID int, productID int) (bool, error)
}

type ProductService interface {
	GetQtyOfProduct(productID int) (int, error)
	LockProduct(productID int, howMany int) error
}

type ProductLocker interface {
	AllocateProductToUser(userID int, productID int, howMany int)
}
