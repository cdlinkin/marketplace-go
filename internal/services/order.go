package services

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(o *models.Order) error {
	if err := o.Validate(); err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return nil
}
func (s *OrderService) CalculatorTotal(o *models.Order) float64 {
	return o.Total()
}
