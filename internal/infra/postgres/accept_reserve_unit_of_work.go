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
	db *gorm.DB,
	ctx context.Context,
) (*AcceptReserveUnitOfWork, error) {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, apperror.Wrap(tx.Error).
			UnExpected().
			Build()
	}
	return &AcceptReserveUnitOfWork{
		tx:              tx,
		inventoryRepo:   NewInventoryRepository(tx),
		transactionRepo: NewTransactionRepository(tx),
	}, nil
}

func (iuw *AcceptReserveUnitOfWork) AcceptReserve(final domain.FinalizeReservation) error {
	return iuw.inventoryRepo.AcceptReserve(final)
}

func (iuw *AcceptReserveUnitOfWork) AcceptHistory(ID string) error {
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
