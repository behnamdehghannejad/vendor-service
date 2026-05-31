package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryUnitOfWork interface {
	IncreaseReserveInventory(domain.RequestReserve) error
	CreateHistory(domain.History) error
	Commit() error
	Rollback() error
}
