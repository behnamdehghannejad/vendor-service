package scheduler

import (
	"sync"
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	productService port.ProductService
	sch            *gocron.Scheduler
}

func New(
	productService port.ProductService,
) *Scheduler {
	return &Scheduler{
		productService: productService,
		sch:            gocron.NewScheduler(time.UTC),
	}
}

func (s *Scheduler) Start(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	s.sch.Cron("*/5 * * * *").Do(s.productService.UpdateAllProductsDiscountPercentage)
	s.sch.StartAsync()

	<-done
	s.sch.Stop()
}
