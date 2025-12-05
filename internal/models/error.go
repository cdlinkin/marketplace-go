package models

import "errors"

var (
	ErrNameEmpty       = errors.New("Error: The name cannot be empty")
	ErrInvalidPrice    = errors.New("Error: The price is invalid")
	ErrInvalidQuantity = errors.New("Error: The quantity is invalid")
	ErrInvalidEmail    = errors.New("Error: The email is invalid")
	ErrInvalidPassword = errors.New("Error: The password is invalid")
)
