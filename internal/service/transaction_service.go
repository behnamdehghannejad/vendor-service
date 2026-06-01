package service

import (
	"context"
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/log"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type HistoryService struct {
	repository        port.HistoryRepository
	unitOfWorkFactory port.UnitOfWorkFactor
}

func NewHistoryService(repository port.HistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func (s *HistoryService) Update(transaction domain.Transaction) error {
	return s.repository.Update(transaction)
}

func (s *HistoryService) Search(search domain.SearchTransaction) ([]domain.Transaction, error) {
	return s.repository.Filter(search)
}

func (s *HistoryService) Approve(ID string) error {
	transaction, err := s.repository.GetByID(ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	auw, err := s.unitOfWorkFactory.AcceptReserveInventoryUnitOfWork(ctx)

	defer func() {
		if err == nil {
			return
		}
		if errRollback := auw.Rollback(); errRollback != nil {
			log.Debugf("error to rollback %v", errRollback)
		}
	}()

	err = auw.AcceptHistory(ID)
	if err != nil {
		return err
	}

	err = auw.AcceptReserve(domain.FinalizeReservation{
		VendorID:  transaction.VendorID,
		ProductID: transaction.ProductID,
		Reserve:   transaction.Reserved,
	})

	err = auw.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *HistoryService) Reject(ID string) error {
	transaction, err := s.repository.GetByID(ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	auw, err := s.unitOfWorkFactory.RejectReserveInventoryUnitOfWork(ctx)

	defer func() {
		if err == nil {
			return
		}
		if errRollback := auw.Rollback(); errRollback != nil {
			log.Debugf("error to rollback %v", errRollback)
		}
	}()

	err = auw.RejectHistory(ID)
	if err != nil {
		return err
	}

	err = auw.RejectReserve(domain.FinalizeReservation{
		VendorID:  transaction.VendorID,
		ProductID: transaction.ProductID,
		Reserve:   transaction.Reserved,
	})

	err = auw.Commit()
	if err != nil {
		return err
	}
	return nil
}
