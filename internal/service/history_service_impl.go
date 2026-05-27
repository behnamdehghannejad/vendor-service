package service

import (
	"github.com/behnamdehghannejad/vendor/internal/domain"
	"github.com/behnamdehghannejad/vendor/internal/infra/repository"

	"github.com/google/uuid"
)

type HistoryServiceImpl struct {
	repository repository.HistoryRepository
}

func NewHistoryService(repository repository.HistoryRepository) *HistoryServiceImpl {
	return &HistoryServiceImpl{repository: repository}
}

func (service *HistoryServiceImpl) Create(history *domain.History) error {
	return service.repository.Add(history)
}

func (service *HistoryServiceImpl) Update(history *domain.History) error {
	return service.repository.Update(history)
}

func (service *HistoryServiceImpl) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *HistoryServiceImpl) FindByOrderID(id uuid.UUID) (*domain.History, error) {
	return service.repository.FindByOrderID(id)
}

func (service *HistoryServiceImpl) FindByPaymentID(paymentID uuid.UUID) (*domain.History, error) {
	return service.repository.FindByPaymentID(paymentID)
}

func (service *HistoryServiceImpl) FindByProductID(productID int) ([]domain.History, error) {
	return service.repository.FindByProductID(productID)
}

func (service *HistoryServiceImpl) FindByVendorID(vendorID int) ([]domain.History, error) {
	return service.repository.FindByVendorID(vendorID)
}

func (service *HistoryServiceImpl) FindByStatus(status domain.Status) ([]domain.History, error) {
	panic("implement me")
}

func (service *HistoryServiceImpl) FindByIsActive(isActive bool) ([]domain.History, error) {
	panic("implement me")
}
