package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cdlinkin/marketplace/internal/api/dto"
	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/services"
)

func orderToResponse(o *models.Order) dto.OrderResponse {
	items := make([]dto.OrderItemResponse, 0, len(o.Items))
	for _, it := range o.Items {
		items = append(items, dto.OrderItemResponse{
			ProductID: it.ProductID,
			Quantity:  it.Quantity,
			Price:     it.Price,
		})
	}

	return dto.OrderResponse{
		ID:        o.ID,
		UserID:    o.UserID,
		Status:    o.Status,
		CreatedAt: o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Items:     items,
		Total:     o.Total(),
	}
}

func CreateOrder(
	orderService *services.OrderService,
	cartService *services.CartService,
	productService *services.ProductService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var req dto.CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		cart, err := cartService.GetCart(req.UserID)
		if err != nil {
			http.Error(w, "cart not found", http.StatusBadRequest)
			return
		}

		if len(cart.Items) == 0 {
			http.Error(w, "cart is empty", http.StatusBadRequest)
			return
		}

		order := &models.Order{
			UserID:    req.UserID,
			Status:    "pending",
			CreatedAt: time.Now(),
		}

		for productID, qty := range cart.Items {
			p, err := productService.GetID(productID)
			if err != nil {
				http.Error(w, "product not found in order build", http.StatusBadRequest)
				return
			}
			order.AddItem(productID, qty, p.Price)
		}

		if err := orderService.CreateOrder(order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := orderToResponse(order)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func GetOrderByID(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		path := r.URL.Path
		prefix := "/order/"

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

		order, err := orderService.GetOrder(id)
		if err != nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}

		resp := orderToResponse(order)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
