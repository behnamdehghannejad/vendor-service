package port

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type InventoryService interface {
	AddProductsToVendor(inventory domain.Inventory) error
	FindByVendorIDAndProductID(vendorID int, productID int) (domain.Inventory, error)
	Update(inventory domain.Inventory) error
}
