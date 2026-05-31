package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryRepository interface {
	Create(domain.Inventory) error
	GetInventory(int, int) (domain.Inventory, error)
}
