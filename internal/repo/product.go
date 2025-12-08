package repo

import "github.com/cdlinkin/marketplace/internal/models"

type ProductRepo interface {
	Save(*models.Product) error
	Get(id int) (*models.Product, error)
	List() ([]models.Product, error)
}

type MemoryProductRepo struct {
	products map[int]*models.Product
	nextID   int
}

func NewMemoryProductRepo() *MemoryProductRepo {
	return &MemoryProductRepo{
		products: make(map[int]*models.Product),
		nextID:   0,
	}
}

func (r *MemoryProductRepo) Save(p *models.Product) error {
	r.nextID++
	p.ID = r.nextID
	r.products[p.ID] = p
	return nil
}

func (r *MemoryProductRepo) Get(id int) (*models.Product, error) {
	p := r.products[id]
	return p, nil
}

func (r *MemoryProductRepo) List() ([]models.Product, error) {
	lists := make([]models.Product, 0, len(r.products))
	for _, p := range r.products {
		lists = append(lists, *p)
	}

	return lists, nil
}
