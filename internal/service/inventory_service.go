package service

import (
	"vendor-service/internal/domain"

	"github.com/google/uuid"
)

type InventoryService interface {
	FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error)
	Update(inventory *domain.Inventory) error
	FindByOrderID(orderID uuid.UUID) (*domain.Inventory, error)
}
