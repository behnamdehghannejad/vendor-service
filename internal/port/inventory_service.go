package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryService interface {
	ReserveQuantity(domain.ReserveRequest) error
	FindInventory(int, int) (domain.Inventory, error)
	Upsert(domain.Inventory) error
	Search(domain.SearchInventory) ([]domain.Inventory, error)
}
