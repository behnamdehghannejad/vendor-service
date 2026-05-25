package repository

type InventoryEntity struct {
	ID int `gorm:"primary_key"`

	VendorID int          `gorm:"column:vendor_id"`
	Vendor   VendorEntity `gorm:"foreignKey:VendorID"`

	ProductID int           `gorm:"column:Product_id"`
	Product   ProductEntity `gorm:"foreignKey:ProductID"`

	Quantity int `gorm:"column:quantity"`
	Reserved int `gorm:"column:reserved"`
}
