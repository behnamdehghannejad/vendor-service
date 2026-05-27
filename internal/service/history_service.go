package service

import (
	"github.com/behnamdehghannejad/vendor/internal/domain"

	"github.com/google/uuid"
)

type HistoryService interface {
	Create(history *domain.History) error
	Update(history *domain.History) error
	Delete(id int) error
	FindByOrderID(id uuid.UUID) (*domain.History, error)
	FindByPaymentID(paymentID uuid.UUID) (*domain.History, error)
	FindByProductID(productID int) ([]domain.History, error)
	FindByVendorID(vendorID int) ([]domain.History, error)
	FindByStatus(status domain.Status) ([]domain.History, error)
	FindByIsActive(isActive bool) ([]domain.History, error)
}
