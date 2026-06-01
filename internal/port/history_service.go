package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryService interface {
	Update(domain.History) error
	Search(domain.SearchHistory) ([]domain.History, error)
}
