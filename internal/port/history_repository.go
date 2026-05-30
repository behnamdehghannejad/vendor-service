package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryRepository interface {
	Add(domain.History) error
	Update(domain.History) error
	FindByOrderID(string) (domain.History, error)
	Filter(domain.SearchHistory) ([]domain.History, error)
}
