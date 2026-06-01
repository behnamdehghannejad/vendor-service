package postgres

import (
	"context"

	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"gorm.io/gorm"
)

type UnitOfWordFactory struct {
	db *gorm.DB
	tx *gorm.DB

	inventoryRepo   *InventoryRepository
	transactionRepo *TransactionRepository
}

func NewUnitOfWorkFactory(db *gorm.DB) *UnitOfWordFactory {
	return &UnitOfWordFactory{
		db: db,
	}
}

func (uof *UnitOfWordFactory) CreateInventoryUnitOfWork(ctx context.Context) (port.ReserveInventoryUnitOfWork, error) {
	return NewReserveInventoryUnitOfWork(uof.db, ctx)
}

func (uof *UnitOfWordFactory) AcceptReserveInventoryUnitOfWork(ctx context.Context) (port.AcceptInventoryUnitOfWork, error) {
	return NewAcceptReserveUnitOfWork(uof.db, ctx)
}

func (uof *UnitOfWordFactory) RejectReserveInventoryUnitOfWork(ctx context.Context) (port.RejectInventoryUnitOfWork, error) {
	return NewRejectReserveUnitOfWork(uof.db, ctx)
}
