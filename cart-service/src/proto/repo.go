package proto

type CartDB interface {
	IsUserHasProduct(userID, productID int) (bool, error)
	ByID(userID int) (Cart, error)
	AddThisToUser(userID int, productID int, buyingStock int) error
}
