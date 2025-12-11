package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cdlinkin/marketplace/internal/api/dto"
	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/services"
)

func cartToResponse(cart *models.Cart, productService *services.ProductService) dto.CartResponse {
	items := make([]dto.CartItemResponse, 0, len(cart.Items))
	prices := make(map[int]float64, len(cart.Items))

	for productID, qty := range cart.Items {
		p, err := productService.GetID(productID)
		if err != nil {
			continue
		}
		prices[productID] = p.Price

		items = append(items, dto.CartItemResponse{
			ProductID: productID,
			Quantity:  qty,
			Price:     p.Price,
		})
	}

	total := cart.Total(prices)

	return dto.CartResponse{
		UserID: cart.UserID,
		Items:  items,
		Total:  total,
	}
}

func AddToCart(cartService *services.CartService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dto.AddToCartRequest
		json.NewDecoder(r.Body).Decode(&req)

		if req.Quantity <= 0 {
			http.Error(w, "Quantity must be > 0", http.StatusBadRequest)
			return
		}
		cartService.AddProduct(req.UserID, req.ProductID, req.Quantity)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}
}

func GetCart(cartService *services.CartService, productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userIDStr := r.URL.Query().Get("user_id")
		if userIDStr == "" {
			http.Error(w, "user_id is required", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "invalid user_id", http.StatusBadRequest)
			return
		}

		cart, err := cartService.GetCart(userID)
		if err != nil {
			http.Error(w, "cart not found", http.StatusNotFound)
			return
		}

		resp := cartToResponse(cart, productService)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
