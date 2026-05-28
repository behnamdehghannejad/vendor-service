package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"

	"github.com/google/uuid"
)

type HistoryService interface {
	Create(domain.History) error
	Update(domain.History) error
	Delete(int) error
	FindByOrderID(uuid.UUID) (domain.History, error)
	FindByPaymentID(uuid.UUID) (domain.History, error)
	FindByProductID(int) ([]domain.History, error)
	FindByVendorID(int) ([]domain.History, error)
	FindByStatus(domain.Status) ([]domain.History, error)
	FindByIsActive(bool) ([]domain.History, error)
}
