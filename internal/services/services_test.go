package services_test

import (
	"testing"

	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/repo"
	"github.com/cdlinkin/marketplace/internal/services"
)

func TestProductService_Create_Simple(t *testing.T) {
	productRepo := repo.NewProductRepo("memory")
	svc := services.NewProductService(productRepo)

	tests := []struct {
		name      string
		product   *models.Product
		wantError bool
	}{
		{
			name: "valid product",
			product: &models.Product{
				Name:     "Phone",
				Price:    100,
				Quantity: 3,
			},
			wantError: false,
		},
		{
			name: "invalid product (price 0)",
			product: &models.Product{
				Name:     "Goat",
				Price:    0,
				Quantity: 10,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Create(tt.product)

			if tt.wantError && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
