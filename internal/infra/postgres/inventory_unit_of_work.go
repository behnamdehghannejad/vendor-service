package postgres

import (
	"context"
	"errors"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"gorm.io/gorm"
)

type InventoryUnitOfWork struct {
	tx *gorm.DB

	inventoryRepo *InventoryRepository
	historyRepo   *HistoryRepository
}

func NewInventoryUnitOfWork(
	db *gorm.DB,
	ctx context.Context,
) (*InventoryUnitOfWork, error) {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, apperror.Wrap(tx.Error).
			UnExpected().
			Build()
	}
	return &InventoryUnitOfWork{
		tx:            tx,
		inventoryRepo: NewInventoryRepository(tx),
		historyRepo:   NewHistoryRepository(tx),
	}, nil
}

func (iuw *InventoryUnitOfWork) IncreaseReserveInventory(vendorID int, productID int, reserved int) error {
	return iuw.inventoryRepo.IncreaseReserveInventory(vendorID, productID, reserved)
}

func (iuw *InventoryUnitOfWork) CreateHistory(history domain.History) error {
	return iuw.historyRepo.Create(history)
}

func (iuw *InventoryUnitOfWork) Commit() error {
	if iuw.tx == nil {
		return apperror.Wrap(errors.New("transaction has not been started")).
			UnExpected().
			Build()
	}
	err := iuw.tx.Commit().Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}

	return nil
}

func (iuw *InventoryUnitOfWork) Rollback() error {
	if iuw.tx == nil {
		return apperror.Wrap(errors.New("transaction has not been started")).
			UnExpected().
			Build()
	}
	err := iuw.tx.Rollback().Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}

	return nil
}
