package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryService interface {
	Create(domain.History) error
	Update(domain.History) error
	FindByOrderID(string) (domain.History, error)
	Search(domain.SearchHistory) ([]domain.History, error)
}
