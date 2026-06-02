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

	inventoryRepo   *InventoryRepository
	transactionRepo *TransactionRepository
}

func NewReserveInventoryUnitOfWork(
	tx *gorm.DB,
	ctx context.Context,
) *ReserveInventoryUnitOfWork {
	return &ReserveInventoryUnitOfWork{
		tx:              tx,
		inventoryRepo:   NewInventoryRepository(tx),
		transactionRepo: NewTransactionRepository(tx),
	}
}

func (iuw *ReserveInventoryUnitOfWork) IncreaseReserveInventory(requestReserve domain.RequestReserve) error {
	return iuw.inventoryRepo.IncreaseReserveInventory(requestReserve)
}

func (iuw *ReserveInventoryUnitOfWork) CreateHistory(transaction domain.Transaction) error {
	return iuw.transactionRepo.Create(transaction)
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
