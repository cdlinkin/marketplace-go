package models

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string
	CreatedAt time.Time `json:"createdat"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameEmpty
	}
	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}
	if len(u.Password) < 6 {
		return ErrInvalidPassword
	}
	return nil
}

func NewUser(name, email, pass string) (*User, error) {
	nu := &User{
		ID:        0, // id++, don`t forget
		Name:      name,
		Email:     email,
		Password:  pass,
		CreatedAt: time.Now(),
	}
	if err := nu.Validate(); err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}
	return nu, nil
}
