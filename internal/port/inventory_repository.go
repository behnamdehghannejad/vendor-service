package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryRepository interface {
	Create(domain.Inventory) error
	FindInventory(int, int) (domain.Inventory, error)
}
