package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/cdlinkin/marketplace/internal/api"
	"github.com/cdlinkin/marketplace/internal/async"
	"github.com/cdlinkin/marketplace/internal/repo"
	"github.com/cdlinkin/marketplace/internal/services"
)

func main() {
	productRepo := repo.NewProductRepo("memory")
	orderRepo := repo.NewOrderRepo("memory")

	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo)
	cartService := services.NewCartService()

	orderPool := async.NewOrderWorkelPool(100, orderService)
	orderPool.Start(3)

	handler := api.NewRouter(productService, cartService, orderService, orderPool)

	server := &http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() { _ = server.ListenAndServe() }()

	<-ctx.Done()
	orderPool.Stop()
	_ = server.Shutdown(context.Background())
}
