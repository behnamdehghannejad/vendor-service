package port

type InventoryService interface {
	ReserveQuantity(vendorID int, productID int, reserved int) error
}
