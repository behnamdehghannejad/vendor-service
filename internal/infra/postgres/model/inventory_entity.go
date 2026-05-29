package model

type InventoryEntity struct {
	ID        int `gorm:"primary_key"`
	VendorID  int `gorm:"column:vendor_id"`
	ProductID int `gorm:"column:product_id"`
	Quantity  int `gorm:"column:quantity"`
	Reserved  int `gorm:"column:reserved"`
}

func (InventoryEntity) TableName() string {
	return "inventory"
}
