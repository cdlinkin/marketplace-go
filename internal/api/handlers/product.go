package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/cdlinkin/marketplace/internal/api/dto"
	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/services"
)

func productToResponse(p *models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func CreateProduct(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var req dto.CreateProductRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		p := &models.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Quantity:    req.Quantity,
		}

		err = productService.Create(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := productToResponse(p)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func ListProduct(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		products, err := productService.List()
		if err != nil {
			http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
			return
		}

		response := make([]dto.ProductResponse, 0, len(products))
		for _, p := range products {
			pCopy := p
			response = append(response, productToResponse(&pCopy))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetProductByID(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		path := r.URL.Path
		prefix := "/products/"

		if !strings.HasPrefix(path, prefix) || len(path) <= len(prefix) {
			http.Error(w, "Invalid path", http.StatusBadRequest)
			return
		}

		idStr := path[len(prefix):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		product, err := productService.GetID(id)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		resp := productToResponse(product)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
