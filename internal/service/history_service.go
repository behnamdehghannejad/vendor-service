package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"

	"github.com/google/uuid"
)

type HistoryServiceImpl struct {
	repository port.HistoryRepository
}

func NewHistoryService(repository port.HistoryRepository) *HistoryServiceImpl {
	return &HistoryServiceImpl{repository: repository}
}

func (s *HistoryServiceImpl) Create(history domain.History) error {
	return s.repository.Add(history)
}

func (s *HistoryServiceImpl) Update(history domain.History) error {
	return s.repository.Update(history)
}

func (s *HistoryServiceImpl) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *HistoryServiceImpl) FindByOrderID(id uuid.UUID) (domain.History, error) {
	return s.repository.FindByOrderID(id)
}

func (s *HistoryServiceImpl) FindByPaymentID(paymentID uuid.UUID) (domain.History, error) {
	return s.repository.FindByPaymentID(paymentID)
}

func (s *HistoryServiceImpl) FindByProductID(productID int) ([]domain.History, error) {
	return s.repository.FindByProductID(productID)
}

func (s *HistoryServiceImpl) FindByVendorID(vendorID int) ([]domain.History, error) {
	return s.repository.FindByVendorID(vendorID)
}

func (s *HistoryServiceImpl) FindByStatus(status domain.Status) ([]domain.History, error) {
	panic("implement me")
}

func (s *HistoryServiceImpl) FindByIsActive(isActive bool) ([]domain.History, error) {
	panic("implement me")
}
