package services

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/repo"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	err := u.repo.Save(user)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return nil
}

func (u *UserService) GetID(id int) (*models.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	return user, nil
}
func (u *UserService) GetEmail(email string) (*models.User, error) {
	user, err := u.repo.GetEmail(email)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	return user, nil
}
