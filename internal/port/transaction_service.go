package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type TransactionService interface {
	Update(domain.Transaction) error
	Search(domain.SearchTransaction) ([]domain.Transaction, error)
}
