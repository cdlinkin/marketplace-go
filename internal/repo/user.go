package repo

import (
	"fmt"

	"github.com/cdlinkin/marketplace/internal/models"
)

type UserRepo interface {
	Save(*models.User) error
	Get(id int) (*models.User, error)
	GetEmail(email string) (*models.User, error)
}

type MemoryUserRepo struct {
	users  map[int]*models.User
	nextID int
}

func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{
		users:  make(map[int]*models.User),
		nextID: 0,
	}
}

func (m *MemoryUserRepo) Save(u *models.User) error {
	m.nextID++
	u.ID = m.nextID
	m.users[u.ID] = u
	return nil
}

func (m *MemoryUserRepo) Get(id int) (*models.User, error) {
	p := m.users[id]
	return p, nil
}
func (m *MemoryUserRepo) GetEmail(email string) (*models.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, fmt.Errorf("Error: user not found")
}
