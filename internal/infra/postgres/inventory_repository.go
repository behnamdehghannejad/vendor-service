package postgres

import (
	"fmt"
	"strings"

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
	inventoryModel := repo.toInventoryModel(inventory)
	err := repo.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "vendor_id"},
			{Name: "product_id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"quantity": gorm.Expr("EXCLUDED.quantity"),
			"price":    gorm.Expr("EXCLUDED.price"),
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

func (repo *InventoryRepository) IncreaseReserveInventory(requestReserve domain.RequestReserve) error {
	err := repo.db.Model(&model.InventoryModel{}).
		Where("product_id = ? AND vendor_id = ?", requestReserve.ProductID, requestReserve.VendorID).
		Updates(map[string]interface{}{
			"reserved":   gorm.Expr("reserved + ?", requestReserve.Reserved),
			"version":    gorm.Expr("version + 1"),
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *InventoryRepository) GetInventory(vendorID int, productID int) (domain.Inventory, error) {
	var inventoryModel model.InventoryModel
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryModel).Error; err != nil {
		return domain.Inventory{}, convertPostgresErrorToAppError(err)
	}

	return repo.toInventoryDomain(inventoryModel), nil
}

func (repo *InventoryRepository) GetInventoryByVendorAndProduct(vendorID int, productID int) (domain.Inventory, error) {
	var inventoryModel model.InventoryModel
	if err := repo.db.Where("vendor_id = ? AND product_id = ?", vendorID, productID).First(&inventoryModel).Error; err != nil {
		return domain.Inventory{}, convertPostgresErrorToAppError(err)
	}

	return repo.toInventoryDomain(inventoryModel), nil
}

func (repo *InventoryRepository) Filter(filter domain.SearchInventory) ([]domain.Inventory, error) {
	var inventories []model.InventoryModel

	query := repo.db.Model(&model.InventoryModel{})

	if filter.VendorID != nil {
		query = query.Where("vendor_id = ?", *filter.VendorID)
	}

	if filter.ProductID != nil {
		query = query.Where("product_id = ?", *filter.ProductID)
	}

	err := query.Find(&inventories).Error
	if err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	return repo.toInventoryDomains(inventories), nil
}

func (repo *InventoryRepository) AcceptReserve(final domain.FinalizeReservation) error {
	err := repo.db.Model(&model.InventoryModel{}).
		Where("vendor_id = ? AND product_id = ?", final.VendorID, final.ProductID).
		Updates(map[string]interface{}{
			"quantity":   gorm.Expr("quantity - ?", final.Reserve),
			"reserved":   gorm.Expr("reserved - ?", final.Reserve),
			"version":    gorm.Expr("version + 1"),
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *InventoryRepository) DeleteInventoriesByID(vendorID int, productID int) error {
	err := repo.db.
		Where("vendor_id = ? AND product_id = ?", vendorID, productID).
		Delete(&model.InventoryModel{}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *InventoryRepository) UpdateProductDiscountPercentages(
	items []domain.InventoryDiscountPercentage,
) error {
	if len(items) == 0 {
		return nil
	}

	var (
		caseStmt strings.Builder
		whereIDs []string
		args     []interface{}
	)

	caseStmt.WriteString("CASE ")

	for _, item := range items {
		caseStmt.WriteString(
			"WHEN vendor_id = ? AND product_id = ? THEN ? ",
		)

		args = append(
			args,
			item.VendorID,
			item.ProductID,
			item.DiscountPercentage,
		)

		whereIDs = append(
			whereIDs,
			fmt.Sprintf("(%d,%d)", item.VendorID, item.ProductID),
		)
	}

	caseStmt.WriteString("ELSE discount_percentage END")

	query := fmt.Sprintf(`
		UPDATE inventories
		SET discount_percentage = %s
		WHERE (vendor_id, product_id) IN (%s)
	`,
		caseStmt.String(),
		strings.Join(whereIDs, ","),
	)

	return repo.db.Exec(query, args...).Error
}

func (repo *InventoryRepository) RejectReserve(final domain.FinalizeReservation) error {
	err := repo.db.Model(&model.InventoryModel{}).
		Where("vendor_id = ? AND product_id = ?", final.VendorID, final.ProductID).
		Updates(map[string]interface{}{
			"reserved":   gorm.Expr("reserved - ?", final.Reserve),
			"version":    gorm.Expr("version + 1"),
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *InventoryRepository) toInventoryDomains(inventoryModels []model.InventoryModel) []domain.Inventory {
	inventoryDomains := make([]domain.Inventory, 0, len(inventoryModels))
	for _, inventoryModel := range inventoryModels {
		inventoryDomains = append(inventoryDomains, repo.toInventoryDomain(inventoryModel))
	}
	return inventoryDomains
}

func (repo *InventoryRepository) toInventoryModel(inventory domain.Inventory) model.InventoryModel {
	return model.InventoryModel{
		ProductID:          inventory.ProductID,
		VendorID:           inventory.VendorID,
		Quantity:           inventory.Quantity,
		Reserved:           inventory.Reserved,
		DiscountPercentage: inventory.DiscountPercentage,
		V:                  inventory.V,
		Price:              inventory.Price,
	}
}

func (repo *InventoryRepository) toInventoryDomain(inventoryModel model.InventoryModel) domain.Inventory {
	return domain.Inventory{
		VendorID:           inventoryModel.VendorID,
		ProductID:          inventoryModel.ProductID,
		Quantity:           inventoryModel.Quantity,
		Reserved:           inventoryModel.Reserved,
		DiscountPercentage: inventoryModel.DiscountPercentage,
		Price:              inventoryModel.Price,
		V:                  inventoryModel.V,
	}
}
