package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryService interface {
	Update(domain.Transaction) error
	Search(domain.SearchTransaction) ([]domain.Transaction, error)
}
