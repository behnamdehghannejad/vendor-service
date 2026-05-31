package postgres

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) Add(product domain.Product) error {
	entity := repo.toProductModel(product)

	if err := repo.db.Create(entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, product)
	}

	return nil
}

func (repo *ProductRepository) Update(product domain.Product) error {
	entity := repo.toProductModel(product)

	if err := repo.db.Save(entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, product)
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	if err := repo.db.Model(&model.ProductModel{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"active":     false,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return convertPostgresErrorToAppError(err, id)
	}

	return nil
}

func (repo *ProductRepository) Filter(filter domain.SearchProduct) ([]domain.Product, error) {
	var entities []model.ProductModel

	query := repo.db.Model(&model.ProductModel{})

	if filter.SearchName != "" {
		query = query.Where(
			"name ILIKE ?",
			"%"+filter.SearchName+"%",
		)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	return repo.toProductsDomain(entities), nil
}

func (repo *ProductRepository) FindById(id int) (domain.Product, error) {
	var entity model.ProductModel

	if err := repo.db.First(&entity, id).Error; err != nil {
		return domain.Product{}, convertPostgresErrorToAppError(err)
	}

	return repo.toProductDomain(entity), nil
}

func (repo *ProductRepository) toProductModel(product domain.Product) *model.ProductModel {
	return &model.ProductModel{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func (repo *ProductRepository) toProductDomain(product model.ProductModel) domain.Product {
	return domain.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func (repo *ProductRepository) toProductsDomain(
	entities []model.ProductModel,
) []domain.Product {
	products := make([]domain.Product, 0, len(entities))

	for _, entity := range entities {
		products = append(
			products,
			repo.toProductDomain(entity),
		)
	}

	return products
}
