package models

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdat"`
}
