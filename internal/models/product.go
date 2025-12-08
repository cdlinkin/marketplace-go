package models

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdat"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrNameEmpty
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	if p.Quantity < 0 {
		return ErrInvalidQuantity
	}

	return nil
}

func NewProduct(name, description string, price float64, quantity int) (*Product, error) {
	np := &Product{
		ID:          0,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   time.Now(),
	}

	if err := np.Validate(); err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	return np, nil
}
