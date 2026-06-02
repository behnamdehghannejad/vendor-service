package postgres

import (
	"context"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
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
	tx, err := uof.createTransactionConnection(ctx)
	if err != nil {
		return nil, err
	}
	return NewReserveInventoryUnitOfWork(tx, ctx), nil
}

func (uof *UnitOfWordFactory) AcceptReserveInventoryUnitOfWork(ctx context.Context) (port.AcceptInventoryUnitOfWork, error) {
	tx, err := uof.createTransactionConnection(ctx)
	if err != nil {
		return nil, err
	}
	return NewAcceptReserveUnitOfWork(tx, ctx), nil
}

func (uof *UnitOfWordFactory) RejectReserveInventoryUnitOfWork(ctx context.Context) (port.RejectInventoryUnitOfWork, error) {
	tx, err := uof.createTransactionConnection(ctx)
	if err != nil {
		return nil, err
	}
	return NewRejectReserveUnitOfWork(tx, ctx), nil
}

func (uof *UnitOfWordFactory) createTransactionConnection(ctx context.Context) (*gorm.DB, error) {
	tx := uof.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, apperror.Wrap(tx.Error).
			UnExpected().
			DebuggingError().
			Log().
			Build()
	}
	return tx, nil
}
