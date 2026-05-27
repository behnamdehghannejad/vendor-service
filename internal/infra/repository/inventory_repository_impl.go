package repository

import (
	"vendor-service/internal/domain"

	"gorm.io/gorm"
)

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

type InventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepositoryImpl(db *gorm.DB) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{
		db: db,
	}
}

func (repo *InventoryRepositoryImpl) Add(inventory *domain.Inventory) error {
	return repo.db.Save(toInventoryEntity(inventory)).Error
}

func (repo *InventoryRepositoryImpl) FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error) {
	var inventoryEntity InventoryEntity
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryEntity).Error; err != nil {
		return nil, err
	}

	return toInventoryDomain(inventoryEntity), nil
}

func (repo *InventoryRepositoryImpl) Update(inventory *domain.Inventory) error {
	return repo.db.Save(toInventoryEntity(inventory)).Error
}
