package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type RejectInventoryUnitOfWork interface {
	RejectReserve(domain.FinalizeReservation) error
	RejectTransaction(string) error
	Commit() error
	Rollback() error
}
