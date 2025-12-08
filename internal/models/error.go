package models

import "errors"

var (
	ErrOrderTotal         = errors.New("Error: The order-total is invalid(total < 0)")
	ErrOrderStatusInvalid = errors.New("Error: The order-status is invalid")
	ErrNameEmpty          = errors.New("Error: The name cannot be empty")
	ErrOrderItemsEmpty    = errors.New("Error: The order.items cannot be empty")
	ErrInvalidPrice       = errors.New("Error: The price is invalid")
	ErrInvalidQuantity    = errors.New("Error: The quantity is invalid")
	ErrInvalidEmail       = errors.New("Error: The email is invalid")
	ErrInvalidPassword    = errors.New("Error: The password is invalid")
)
