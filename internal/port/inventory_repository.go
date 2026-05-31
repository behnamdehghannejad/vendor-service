package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryRepository interface {
	Create(domain.Inventory) error
	GetInventory(int) (domain.Inventory, error)
	GetInventoryByVendorAndProduct(int, int) (domain.Inventory, error)
}
