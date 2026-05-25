package domain

import "time"

type VendorInventory struct {
	ID       int
	VendorID int
	Vendor   Vendor `gorm:"foreignKey:VendorID"`

	ProductID int
	Product   Product `gorm:"foreignKey:ProductID"`

	Quantity int
	Reserved int

	CreatedAt time.Time
	UpdatedAt time.Time
}
