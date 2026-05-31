package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryService interface {
	ReserveQuantity(int, int, int, string) error
	FindInventory(int, int) (domain.Inventory, error)
	Upsert(domain.Inventory) error
}
