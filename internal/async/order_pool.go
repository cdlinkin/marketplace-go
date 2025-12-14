package async

import (
	"sync"

	"github.com/cdlinkin/marketplace/internal/models"
	"github.com/cdlinkin/marketplace/internal/services"
)

const Complete = "complete"

type OrderJob struct {
	Order *models.Order
}

type OrderWorkerPool struct {
	Jobs chan OrderJob
	wg   sync.WaitGroup
	svc  *services.OrderService
}

func NewOrderWorkelPool(buffer int, svc *services.OrderService) *OrderWorkerPool {
	return &OrderWorkerPool{
		Jobs: make(chan OrderJob, buffer),
		svc:  svc,
	}
}

func (p *OrderWorkerPool) Start(n int) {
	for i := 0; i < n; i++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			for job := range p.Jobs {
				job.Order.Status = Complete
				_ = p.svc.CreateOrder(job.Order)
			}
		}(i + 1)
	}
}

func (p *OrderWorkerPool) Stop() {
	close(p.Jobs)
	p.wg.Wait()
}
