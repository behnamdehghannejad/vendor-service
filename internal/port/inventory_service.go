package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type InventoryService interface {
	ReserveQuantity(vendorID int, productID int, reserved int) error
	FindInventory(int) (domain.Inventory, error)
}
