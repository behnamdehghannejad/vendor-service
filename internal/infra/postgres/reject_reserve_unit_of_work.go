package postgres

import (
	"context"
	"errors"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"gorm.io/gorm"
)

type RejectReserveUnitOfWork struct {
	tx *gorm.DB

	inventoryRepo   *InventoryRepository
	transactionRepo *TransactionRepository
}

func NewRejectReserveUnitOfWork(
	db *gorm.DB,
	ctx context.Context,
) (*RejectReserveUnitOfWork, error) {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, apperror.Wrap(tx.Error).
			UnExpected().
			Build()
	}
	return &RejectReserveUnitOfWork{
		tx:              tx,
		inventoryRepo:   NewInventoryRepository(tx),
		transactionRepo: NewTransactionRepository(tx),
	}, nil
}

func (iuw *RejectReserveUnitOfWork) RejectReserve(final domain.FinalizeReservation) error {
	return iuw.inventoryRepo.AcceptReserve(final)
}

func (iuw *RejectReserveUnitOfWork) RejectHistory(ID string) error {
	return iuw.transactionRepo.Approve(ID)
}

func (iuw *RejectReserveUnitOfWork) Commit() error {
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

func (iuw *RejectReserveUnitOfWork) Rollback() error {
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
