package models

type Cart struct {
	UserID int         `json:"user_id"`
	Items  map[int]int `json:"items"`
}

func (c *Cart) AddProduct(productID, quantity int) {
	c.Items[productID] = quantity
}

func (c *Cart) RemoveProduct(productID int) {
	delete(c.Items, productID)
}

func (c *Cart) Total(productPrice map[int]float64) float64 {
	sum := 0.0
	for prID, quantity := range c.Items {
		sum += productPrice[prID] * float64(quantity)
	}
	return sum
}
