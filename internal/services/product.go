package services

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) Create(p *models.Product) error {
	if err := p.Validate(); err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return nil
}

func (s *ProductService) List() ([]models.Product, error) {
	return []models.Product{}, nil
}

func (s *ProductService) GetID(id int) (*models.Product, error) {
	return nil, nil
}
