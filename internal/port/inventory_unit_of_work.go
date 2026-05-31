package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryUnitOfWork interface {
	IncreaseReserveInventory(int, int, int) error
	CreateHistory(domain.History) error
	Commit() error
	Rollback() error
}
