package repository

import (
	"vendor-service/internal/domain"
)

type InventoryRepository interface {
	Add(inventory *domain.Inventory) error
	FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error)
	Update(inventory *domain.Inventory) error
}
