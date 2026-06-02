package scheduler

import (
	"sync"
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	inventoryService port.InventoryService
	sch              *gocron.Scheduler
}

func New(
	inventoryService port.InventoryService,
) *Scheduler {
	return &Scheduler{
		inventoryService: inventoryService,
		sch:              gocron.NewScheduler(time.UTC),
	}
}

func (s *Scheduler) Start(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	s.sch.Cron("*/5 * * * *").Do(s.inventoryService.UpdateAllInventoriesDiscountPercentage)
	s.sch.StartAsync()

	<-done
	s.sch.Stop()
}
