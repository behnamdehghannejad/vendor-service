package postgres

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (repo *InventoryRepository) Create(inventory domain.Inventory) error {
	return repo.db.Save(toInventoryEntity(inventory)).Error
}

func (repo *InventoryRepository) FindByVendorIDAndProductID(vendorID int, productID int) (domain.Inventory, error) {
	var inventoryEntity model.InventoryModel
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryEntity).Error; err != nil {
		return domain.Inventory{}, err
	}

	return toInventoryDomain(inventoryEntity), nil
}

func (repo *InventoryRepository) Update(inventory domain.Inventory) error {
	return repo.db.Save(toInventoryEntity(inventory)).Error
}

func toInventoryEntity(inventory domain.Inventory) model.InventoryModel {
	return model.InventoryModel{
		ID:        inventory.ID,
		ProductID: inventory.ProductID,
		VendorID:  inventory.VendorID,
		Quantity:  inventory.Quantity,
		Reserved:  inventory.Reserved,
	}
}

func toInventoryDomain(entity model.InventoryModel) domain.Inventory {
	return domain.Inventory{
		ID:        entity.ID,
		VendorID:  entity.VendorID,
		ProductID: entity.ProductID,
		Quantity:  entity.Quantity,
		Reserved:  entity.Reserved,
	}
}
