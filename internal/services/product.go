package services

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/repo"
)

type ProductService struct {
	repo repo.ProductRepo
}

func NewProductService(repo repo.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(p *models.Product) error {
	if err := p.Validate(); err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	s.repo.Save(p)
	return nil
}

func (s *ProductService) List() ([]models.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetID(id int) (*models.Product, error) {
	product, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
