package repo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cdlinkin/marketplace/internal/models"
)

type FileProductRepo struct {
	filename string
}

func NewFileProductRepo(filename string) *FileProductRepo {
	return &FileProductRepo{
		filename: filename,
	}
}

func (f *FileProductRepo) Save(p *models.Product) error {
	products, _ := f.List()
	products = append(products, *p)

	data, _ := json.MarshalIndent(products, "", "  ")
	return os.WriteFile(f.filename, data, 0644)
}

func (f *FileProductRepo) List() ([]models.Product, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	var products []models.Product
	json.Unmarshal(data, &products)
	return products, nil
}

func (f *FileProductRepo) GetID(id int) (*models.Product, error) {
	products, err := f.List()
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("Error: product %d not found", id)
}
