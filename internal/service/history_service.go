package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"

	"github.com/google/uuid"
)

type HistoryService struct {
	repository port.HistoryRepository
}

func NewHistoryService(repository port.HistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func (s *HistoryService) Create(history domain.History) error {
	return s.repository.Add(history)
}

func (s *HistoryService) Update(history domain.History) error {
	return s.repository.Update(history)
}

func (s *HistoryService) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *HistoryService) FindByOrderID(id uuid.UUID) (domain.History, error) {
	return s.repository.FindByOrderID(id)
}

func (s *HistoryService) FindByPaymentID(paymentID uuid.UUID) (domain.History, error) {
	return s.repository.FindByPaymentID(paymentID)
}

func (s *HistoryService) FindByProductID(productID int) ([]domain.History, error) {
	return s.repository.FindByProductID(productID)
}

func (s *HistoryService) FindByVendorID(vendorID int) ([]domain.History, error) {
	return s.repository.FindByVendorID(vendorID)
}

func (s *HistoryService) FindByStatus(status domain.Status) ([]domain.History, error) {
	panic("implement me")
}

func (s *HistoryService) FindByIsActive(isActive bool) ([]domain.History, error) {
	panic("implement me")
}
