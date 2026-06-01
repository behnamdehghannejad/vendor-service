package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type AcceptInventoryUnitOfWork interface {
	AcceptReserve(domain.FinalizeReservation) error
	AcceptHistory(string) error
	Commit() error
	Rollback() error
}
