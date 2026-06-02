package model

type InventoryModel struct {
	VendorID           int     `gorm:"column:vendor_id"`
	ProductID          int     `gorm:"column:product_id"`
	Quantity           int     `gorm:"column:quantity"`
	DiscountPercentage float64 `gorm:"column:discount_percentage"`
	Reserved           int     `gorm:"column:reserved"`
	V                  int     `gorm:"column:version"`
}

func (InventoryModel) TableName() string {
	return "inventories"
}
