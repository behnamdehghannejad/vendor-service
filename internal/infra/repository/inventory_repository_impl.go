package repository

import (
	"vendor-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryEntity struct {
	ID        int           `gorm:"primary_key"`
	VendorID  int           `gorm:"column:vendor_id"`
	Vendor    VendorEntity  `gorm:"foreignKey:VendorID"`
	ProductID int           `gorm:"column:Product_id"`
	Product   ProductEntity `gorm:"foreignKey:ProductID"`
	Quantity  int           `gorm:"column:quantity"`
	Reserved  int           `gorm:"column:reserved"`
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

func (repo *InventoryRepositoryImpl) FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error) {
	var inventoryEntity InventoryEntity
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(inventoryEntity).Error; err != nil {
		return nil, err
	}

	return toInventoryDomain(inventoryEntity), nil
}

func (repo *InventoryRepositoryImpl) Update(inventory *domain.Inventory) error {
	return repo.db.Save(inventory).Error
}

func (repo *InventoryRepositoryImpl) FindByOrderID(orderID uuid.UUID) (*domain.Inventory, error) {
	var inventory InventoryEntity
	err := repo.db.Where("order_id = ?", orderID).Find(&inventory).Error
	if err != nil {
		return nil, err
	}
	return toInventoryDomain(inventory), nil
}

func toInventoryDomain(entity InventoryEntity) *domain.Inventory {
	return &domain.Inventory{
		ID:       entity.ID,
		Vendor:   domain.Vendor(entity.Vendor),
		Product:  domain.Product(entity.Product),
		Quantity: entity.Quantity,
		Reserved: entity.Reserved,
	}
}
