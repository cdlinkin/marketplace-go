package dto

type AddToCartRequest struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartItemResponse struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"prict"`
}

type CartResponse struct {
	UserID int                `json:"user_id"`
	Items  []CartItemResponse `json:"items"`
	Total  float64            `json:"total"`
}
