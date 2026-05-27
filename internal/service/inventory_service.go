package service

import (
	"vendor-service/internal/domain"
)

type InventoryService interface {
	AddProductsToVendor(inventory *domain.Inventory) error
	FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error)
	Update(inventory *domain.Inventory) error
}
