package repo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cdlinkin/marketplace/internal/models"
)

type FileUserRepo struct {
	filename string
}

func NewFileUserRepo(filename string) *FileUserRepo {
	return &FileUserRepo{filename: filename}
}

func (f *FileUserRepo) Save(u *models.User) error {
	users, _ := f.list()
	users = append(users, *u)

	data, _ := json.MarshalIndent(users, "", "  ")
	return os.WriteFile(f.filename, data, 0644)
}

func (f *FileUserRepo) Get(id int) (*models.User, error) {
	users, _ := f.list()

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("Error: user %d not found", id)
}
func (f *FileUserRepo) GetEmail(email string) (*models.User, error) {
	users, _ := f.list()

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("Error: user.email %s not found", email)
}

func (f *FileUserRepo) list() ([]models.User, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	var users []models.User
	json.Unmarshal(data, &users)
	return users, nil
}
