package domain

import "errors"

type Inventory struct {
	ID       int
	Vendor   Vendor
	Product  Product
	Quantity int
	Reserved int
}

func ValidateAndSetQuantity(inventory *Inventory, quantity int) error {
	if quantity > inventory.Quantity {
		return errors.New("inventory quantity is more than the order quantity")
	}

	inventory.Quantity -= quantity
	return nil
}
