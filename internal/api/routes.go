package api

import (
	"net/http"

	"github.com/cdlinkin/marketplace/internal/api/handlers"
	"github.com/cdlinkin/marketplace/internal/async"
	"github.com/cdlinkin/marketplace/internal/services"
)

func NewRouter(
	productService *services.ProductService,
	cartService *services.CartService,
	orderService *services.OrderService,
	orderPool *async.OrderWorkerPool,
) http.HandlerFunc {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateProduct(productService)(w, r)
		case http.MethodGet:
			handlers.ListProduct(productService)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/products/", handlers.GetProductByID(productService))

	mux.HandleFunc("/cart/add", handlers.AddToCart(cartService))
	mux.HandleFunc("/cart", handlers.GetCart(cartService, productService))

	mux.HandleFunc("/order", handlers.CreateOrder(orderService, cartService, productService, orderPool.Jobs))
	mux.HandleFunc("/order/", handlers.GetOrderByID(orderService))

	return mux.ServeHTTP
}
