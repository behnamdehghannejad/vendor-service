package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryRepository interface {
	Upsert(domain.Inventory) error
	GetInventory(int, int) (domain.Inventory, error)
	Filter(domain.SearchInventory) ([]domain.Inventory, error)
	AcceptReserve(domain.FinalizeReservation) error
	RejectReserve(domain.FinalizeReservation) error
}
