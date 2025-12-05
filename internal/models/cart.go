package models

type Cart struct {
	UserID int         `json:"user_id"`
	Items  map[int]int `json:"items"`
}
