package repo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cdlinkin/marketplace/internal/models"
)

type FileOrderRepo struct {
	filename string
}

func NewFileOrderRepo(filename string) *FileOrderRepo {
	return &FileOrderRepo{filename: filename}
}

func (f *FileOrderRepo) Save(o *models.Order) error {
	orders, _ := f.list()
	orders = append(orders, *o)

	data, _ := json.MarshalIndent(orders, "", "  ")
	return os.WriteFile(f.filename, data, 0644)
}

func (f *FileOrderRepo) Get(id int) (*models.Order, error) {
	orders, err := f.list()
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	for _, order := range orders {
		if order.ID == id {
			return &order, nil
		}
	}
	return nil, fmt.Errorf("Error: order %d not found", id)
}

func (f *FileOrderRepo) ListByUser(userID int) ([]models.Order, error) {
	orders, err := f.list()
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	var listOrderByUser []models.Order
	for _, order := range orders {
		if order.UserID == userID {
			listOrderByUser = append(listOrderByUser, order)
		}
	}

	return listOrderByUser, nil
}

func (f *FileOrderRepo) list() ([]models.Order, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	var orders []models.Order
	json.Unmarshal(data, &orders)
	return orders, nil
}
