package services

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/repo"
)

type OrderService struct {
	repo repo.OrderRepo
}

func NewOrderService(repo repo.OrderRepo) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(o *models.Order) error {
	if err := o.Validate(); err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	s.repo.Save(o)

	return nil
}

func (s *OrderService) GetOrder(id int) (*models.Order, error) {
	o, err := s.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	return o, nil
}
func (s *OrderService) ListOrdersByUser(userID int) ([]models.Order, error) {
	list, err := s.repo.ListByUser(userID)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	return list, nil
}
