package postgres

import (
	"context"
	"errors"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"gorm.io/gorm"
)

type ReserveInventoryUnitOfWork struct {
	tx *gorm.DB

	inventoryRepo *InventoryRepository
	historyRepo   *HistoryRepository
}

func NewReserveInventoryUnitOfWork(
	db *gorm.DB,
	ctx context.Context,
) (*ReserveInventoryUnitOfWork, error) {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, apperror.Wrap(tx.Error).
			UnExpected().
			Build()
	}
	return &ReserveInventoryUnitOfWork{
		tx:            tx,
		inventoryRepo: NewInventoryRepository(tx),
		historyRepo:   NewHistoryRepository(tx),
	}, nil
}

func (iuw *ReserveInventoryUnitOfWork) IncreaseReserveInventory(requestReserve domain.RequestReserve) error {
	return iuw.inventoryRepo.IncreaseReserveInventory(requestReserve)
}

func (iuw *ReserveInventoryUnitOfWork) CreateHistory(history domain.History) error {
	return iuw.historyRepo.Create(history)
}

func (iuw *ReserveInventoryUnitOfWork) Commit() error {
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

func (iuw *ReserveInventoryUnitOfWork) Rollback() error {
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
