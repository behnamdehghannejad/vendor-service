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

func (s *HistoryService) Update(history domain.History) error {
	return s.repository.Update(history)
}

func (s *HistoryService) Search(search domain.SearchHistory) ([]domain.History, error) {
	return s.repository.Filter(search)
}

func (s *HistoryService) Approve(ID string) error {
	history, err := s.repository.GetByID(ID)
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
		VendorID:  history.VendorID,
		ProductID: history.ProductID,
		Reserve:   history.Reserved,
	})

	err = auw.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *HistoryService) Reject(ID string) error {
	history, err := s.repository.GetByID(ID)
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
		VendorID:  history.VendorID,
		ProductID: history.ProductID,
		Reserve:   history.Reserved,
	})

	err = auw.Commit()
	if err != nil {
		return err
	}
	return nil
}
