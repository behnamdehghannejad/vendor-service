package model

type InventoryModel struct {
	VendorID  int `gorm:"column:vendor_id"`
	ProductID int `gorm:"column:product_id"`
	Quantity  int `gorm:"column:quantity"`
	Reserved  int `gorm:"column:reserved"`
}

func (InventoryModel) TableName() string {
	return "inventories"
}
