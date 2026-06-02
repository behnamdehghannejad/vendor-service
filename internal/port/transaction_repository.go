package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type TransactionRepository interface {
	Create(domain.Transaction) error
	Update(domain.Transaction) error
	Filter(domain.SearchTransaction) ([]domain.Transaction, error)
	Approve(string) error
	Reject(string) error
	GetByID(string) (domain.Transaction, error)
}
