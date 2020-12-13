package proto

type CartInfoResponse struct {
	UserID int  `json:"userId"`
	Cart   Cart `json:"cart"`
}

type Cart struct {
	Products map[int]*Product `json:"products"`
}

type Product struct {
	ID       int `json:"id"`
	Quantity int `json:"qty"`
}

type AddProductToCardResponse struct {
	UserID    int  `json:"userId"`
	ProductID int  `json:"productId"`
	Success   bool `json:"success"`
}

type UserProduct struct {
	UserID     int  `json:"userId"`
	ProductID  int  `json:"productId"`
	CartStatus bool `json:"cart_status"`
}
