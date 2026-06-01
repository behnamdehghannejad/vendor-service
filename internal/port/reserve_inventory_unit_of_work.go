package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type ReserveInventoryUnitOfWork interface {
	IncreaseReserveInventory(domain.RequestReserve) error
	CreateHistory(domain.Transaction) error
	Commit() error
	Rollback() error
}
