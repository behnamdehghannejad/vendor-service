package postgres

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"

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

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (repo *InventoryRepository) Add(inventory domain.Inventory) error {
	return repo.db.Save(toInventoryEntity(inventory)).Error
}

func (repo *InventoryRepository) FindByVendorIDAndProductID(vendorID int, productID int) (domain.Inventory, error) {
	var inventoryEntity InventoryEntity
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryEntity).Error; err != nil {
		return domain.Inventory{}, err
	}

	return toInventoryDomain(inventoryEntity), nil
}

func (repo *InventoryRepository) Update(inventory domain.Inventory) error {
	return repo.db.Save(toInventoryEntity(inventory)).Error
}

func toInventoryEntity(inventory domain.Inventory) InventoryEntity {
	return InventoryEntity{
		ID:        inventory.ID,
		ProductID: inventory.ProductID,
		VendorID:  inventory.VendorID,
		Quantity:  inventory.Quantity,
		Reserved:  inventory.Reserved,
	}
}

func toInventoryDomain(entity InventoryEntity) domain.Inventory {
	return domain.Inventory{
		ID:        entity.ID,
		VendorID:  entity.VendorID,
		ProductID: entity.ProductID,
		Quantity:  entity.Quantity,
		Reserved:  entity.Reserved,
	}
}
