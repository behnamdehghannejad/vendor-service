package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type RejectInventoryUnitOfWork interface {
	RejectReserve(domain.FinalizeReservation) error
	RejectHistory(string) error
	Commit() error
	Rollback() error
}
