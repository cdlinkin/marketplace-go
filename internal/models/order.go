package models

import "time"

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	Status    string      `json:"status"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"createdat"`
}

func (o *Order) AddItem(productID, quantity int, price float64) {
	newOrderItem := OrderItem{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
	o.Items = append(o.Items, newOrderItem)
}

func (o *Order) Total() float64 {
	sum := 0.0
	for _, v := range o.Items {
		sum += v.Price * float64(v.Quantity)
	}
	return sum
}

func (o *Order) Validate() error {
	if len(o.Items) == 0 {
		return ErrOrderItemsEmpty
	}
	if o.Total() == 0 {
		return ErrOrderTotal
	}

	if o.Status != "pending" && o.Status != "complete" {
		return ErrOrderStatusInvalid
	}

	return nil
}
