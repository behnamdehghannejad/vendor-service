package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryRepository interface {
	Create(domain.History) error
	Update(domain.History) error
	Filter(domain.SearchHistory) ([]domain.History, error)
}
