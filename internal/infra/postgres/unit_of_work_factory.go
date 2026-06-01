package postgres

import (
	"context"

	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"gorm.io/gorm"
)

type UnitOfWordFactory struct {
	db *gorm.DB
	tx *gorm.DB

	inventoryRepo *InventoryRepository
	historyRepo   *HistoryRepository
}

func NewUnitOfWorkFactory(db *gorm.DB) *UnitOfWordFactory {
	return &UnitOfWordFactory{
		db: db,
	}
}

func (uof *UnitOfWordFactory) CreateInventoryUnitOfWork(ctx context.Context) (port.InventoryUnitOfWork, error) {
	return NewReserveInventoryUnitOfWork(uof.db, ctx)
}
