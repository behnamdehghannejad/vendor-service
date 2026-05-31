package postgres

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (repo *InventoryRepository) Upsert(inventory domain.Inventory) error {
	inventoryModel := repo.toInventoryEntity(inventory)
	err := repo.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "vendor_id"},
			{Name: "product_id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"quantity": gorm.Expr("EXCLUDED.quantity"),
			"version":  gorm.Expr("inventories.version + 1"),
		}),
		Where: clause.Where{
			Exprs: []clause.Expression{
				gorm.Expr("inventories.version = ?", inventory.V),
			},
		},
	}).Create(&inventoryModel).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *InventoryRepository) IncreaseReserveInventory(vendorID int, productID int, reserve int) error {
	err := repo.db.Model(&model.InventoryModel{}).
		Where("product_id = ? AND vendor_id = ?", productID, vendorID).
		UpdateColumn("reserved", gorm.Expr("reserved + ?", reserve)).
		Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *InventoryRepository) GetInventory(vendorID int, productID int) (domain.Inventory, error) {
	var inventoryEntity model.InventoryModel
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryEntity).Error; err != nil {
		return domain.Inventory{}, convertPostgresErrorToAppError(err)
	}

	return repo.toInventoryDomain(inventoryEntity), nil
}

func (repo *InventoryRepository) GetInventoryByVendorAndProduct(vendorID int, productID int) (domain.Inventory, error) {
	var inventoryEntity model.InventoryModel
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryEntity).Error; err != nil {
		return domain.Inventory{}, convertPostgresErrorToAppError(err)
	}

	return repo.toInventoryDomain(inventoryEntity), nil
}

func (repo *InventoryRepository) Update(inventory domain.Inventory) error {
	return repo.db.Save(repo.toInventoryEntity(inventory)).Error
}

func (repo *InventoryRepository) toInventoryEntity(inventory domain.Inventory) model.InventoryModel {
	return model.InventoryModel{
		ProductID: inventory.ProductID,
		VendorID:  inventory.VendorID,
		Quantity:  inventory.Quantity,
		Reserved:  inventory.Reserved,
		V:         inventory.V,
	}
}

func (repo *InventoryRepository) toInventoryDomain(entity model.InventoryModel) domain.Inventory {
	return domain.Inventory{
		VendorID:  entity.VendorID,
		ProductID: entity.ProductID,
		Quantity:  entity.Quantity,
		Reserved:  entity.Reserved,
		V:         entity.V,
	}
}
