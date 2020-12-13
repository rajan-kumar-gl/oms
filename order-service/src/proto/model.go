package proto

type GetQuantityResponse struct {
	Data struct {
		ProductID int `json:"id"`
		Qty       int `json:"qty"`
	} `json:"data"`
}

type SetQuantityResponse struct {
	Data struct {
		ProductID int  `json:"id"`
		Success   bool `json:"success"`
	} `json:"data"`
}

type CartServiceResponse struct {
	Data struct {
		UserID     int  `json:"user_id"`
		ProductID  int  `json:"product_id"`
		CartStatus bool `json:"cart_status"`
	} `json:"data"`
}

type OrderInfo struct {
	UserID      int  `json:"user_id"`
	ProductID   int  `json:"product_id"`
	BuyingStock int  `json:"qty"`
	Success     bool `json:"success"`
}
