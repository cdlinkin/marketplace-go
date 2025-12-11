package services

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
)

type CartService struct {
	carts map[int]*models.Cart
}

func NewCartService() *CartService {
	return &CartService{
		carts: make(map[int]*models.Cart),
	}
}

func (s *CartService) AddProduct(userID, productID, quantity int) error {
	cart, err := s.GetCart(userID)
	if err != nil {
		cart = &models.Cart{
			UserID: userID,
			Items:  make(map[int]int),
		}
		s.carts[userID] = cart
	}

	cart.AddProduct(productID, quantity)
	return nil
}

func (s *CartService) GetCart(userID int) (*models.Cart, error) {
	for _, cart := range s.carts {
		if cart.UserID == userID {
			return cart, nil
		}
	}
	return nil, fmt.Errorf("Error: Cart is not founded")
}
