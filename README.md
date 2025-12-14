# Marketplace API

A learning backend project in Go. 
A simple REST API for a mini marketplace with products, a shopping cart, and orders.

The project is made for practice:
- REST API
- backend application architecture
- working with goroutines and channels

## Features

- Products (creation, list retrieval, retrieval by ID)
- User's shopping cart
- Order creation
- Asynchronous order processing via worker pool
- Middleware for logging HTTP requests
- Layer separation: handler → service → repository
- In-memory and file-based repositories (JSON)


## Stack

- Go
- net/http
- JSON
- Middleware (logging)
- Async (Goroutines and channels)
- Worker pool
- Unit-tests

## Architecture

cmd/marketplace/main.go # launching the HTTP server

```
internal/
├── api/ # HTTP handlers and router
├── services/ # business logic
├── repo/ # working with data (memory / file)
├── models/ # domain models
└── async/ # worker pool for orders
```

## Run

```
go mod tidy
go run ./cmd/server
```

The server will start on: ```http://localhost:9090```

## Main endpoints

### Products
- GET /products — list of products
- GET /products/{id} — product by ID
- POST /products — create a product

Create Product - Request
```json
{
  "name": "Phone",
  "description": "I love Golang;)",
  "price": 799,
  "quantity": 5
}
```
Product - Response
```json
{
  "id": 1,
  "name": "Phone",
  "description": "I love Golang;)",
  "price": 799,
  "quantity": 5,
  "created_at": "2025-01-01T12:00:00Z"
}
```

### Cart
- POST /cart/add — add a product to the cart
- GET /cart?user_id=1 — get the user's cart

Add To Cart - Request
```json
{
  "user_id": 1,
  "product_id": 1,
  "quantity": 2
}
```

Cart Item - Response
```json
{
  "user_id": 1,
  "product_id": 1,
  "quantity": 2
}
```

Cart - Response
```json
{
  "user_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "price": 1000
    }
  ],
  "total": 2000
}
```

### Orders
- POST /order — create an order (processed async)
- GET /order/{id} — get an order by ID

Create Order – Request
```json
{
  "user_id": 1
}
```

Order Item – Response
```json
{
  "product_id": 1,
  "quantity": 2,
  "price": 1000
}
```

Order - Response
```json
{
  "id": 1,
  "user_id": 1,
  "status": "pending",
  "created_at": "2025-01-01T12:00:00Z",
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "price": 1000
    }
  ],
  "total": 2000
}
```

## Testing

```
go test ./...
```
