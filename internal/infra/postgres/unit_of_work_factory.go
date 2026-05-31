package postgres

import (
	"context"

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

func (uof *UnitOfWordFactory) CreateInventoryUnitOfWork(ctx context.Context) (*InventoryUnitOfWork, error) {
	return NewInventoryUnitOfWork(uof.db, ctx)
}
