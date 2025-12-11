package dto

type CreateOrderRequest struct {
	UserID int `json:"user_id"`
}

type OrderItemResponse struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderResponse struct {
	ID        int                 `json:"id"`
	UserID    int                 `json:"user_id"`
	Status    string              `json:"status"`
	CreatedAt string              `json:"created_at"`
	Items     []OrderItemResponse `json:"items"`
	Total     float64             `json:"total"`
}
