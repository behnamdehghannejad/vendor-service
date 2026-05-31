package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type InventoryRepository interface {
	Create(domain.Inventory) error
	FindByVendorIDAndProductID(int, int) (domain.Inventory, error)
	Update(domain.Inventory) error
}
