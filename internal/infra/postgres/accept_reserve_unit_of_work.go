package postgres

import (
	"context"
	"errors"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"gorm.io/gorm"
)

type AcceptReserveUnitOfWork struct {
	tx *gorm.DB

	inventoryRepo   *InventoryRepository
	transactionRepo *TransactionRepository
}

func NewAcceptReserveUnitOfWork(
	tx *gorm.DB,
	ctx context.Context,
) *AcceptReserveUnitOfWork {
	return &AcceptReserveUnitOfWork{
		tx:              tx,
		inventoryRepo:   NewInventoryRepository(tx),
		transactionRepo: NewTransactionRepository(tx),
	}
}

func (iuw *AcceptReserveUnitOfWork) AcceptReserve(final domain.FinalizeReservation) error {
	return iuw.inventoryRepo.AcceptReserve(final)
}

func (iuw *AcceptReserveUnitOfWork) AcceptTransaction(ID string) error {
	return iuw.transactionRepo.Approve(ID)
}

func (iuw *AcceptReserveUnitOfWork) Commit() error {
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

func (iuw *AcceptReserveUnitOfWork) Rollback() error {
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
