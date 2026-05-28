package domain

import "time"

type VendorInventory struct {
	ID        int
	VendorID  int
	Vendor    Vendor
	ProductID int
	Product   Product
	Quantity  int
	Reserved  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
