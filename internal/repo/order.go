package repo

import "github.com/cdlinkin/marketplace/internal/models"

type OrderRepo interface {
	Save(*models.Order) error
	Get(id int) (*models.Order, error)
	ListByUser(userID int) ([]models.Order, error)
}

type MemoryOrderRepo struct {
	orders map[int]*models.Order
	nextID int
}

func NewMemoryOrderRepo() *MemoryOrderRepo {
	return &MemoryOrderRepo{
		orders: make(map[int]*models.Order),
		nextID: 0,
	}
}

func (m *MemoryOrderRepo) Save(o *models.Order) error {
	m.nextID++
	o.ID = m.nextID
	m.orders[o.ID] = o
	return nil
}

func (m *MemoryOrderRepo) Get(id int) (*models.Order, error) {
	p := m.orders[id]
	return p, nil
}

func (m *MemoryOrderRepo) ListByUser(userID int) ([]models.Order, error) {
	lists := make([]models.Order, 0)
	for _, o := range m.orders {
		if o.UserID == userID {
			lists = append(lists, *o)
		}
	}

	return lists, nil
}
