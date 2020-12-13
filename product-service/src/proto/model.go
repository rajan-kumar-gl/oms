package proto

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Amt      int    `json:"amt"`
	AvalQty  int    `json:"qty"`
	Currency string `json:"currency"`
}

type GetProductQuantityResponse struct {
	ProductID int `json:"id"`
	AvalQty   int `json:"qty"`
}

type SetProductQuantityResponse struct {
	ProductID int  `json:"id"`
	Success   bool `json:"success"`
}
