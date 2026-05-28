package postgres

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"gorm.io/gorm"
)

type ProductEntity struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"size:255"`
	Description string    `gorm:"size:255"`
	Active      bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (ProductEntity) TableName() string {
	return "products"
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) Add(product domain.Product) error {
	entity := repo.toProductEntity(product)

	if err := repo.db.Create(entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, product)
	}

	return nil
}

func (repo *ProductRepository) Update(product domain.Product) error {
	entity := repo.toProductEntity(product)

	if err := repo.db.Save(entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, product)
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	if err := repo.db.Model(&ProductEntity{}).
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
	var entities []ProductEntity

	query := repo.db.Model(&ProductEntity{})

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
	var entity ProductEntity

	if err := repo.db.First(&entity, id).Error; err != nil {
		return domain.Product{}, convertPostgresErrorToAppError(err)
	}

	return repo.toProductDomain(entity), nil
}

func (repo *ProductRepository) toProductEntity(product domain.Product) *ProductEntity {
	return &ProductEntity{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func (repo *ProductRepository) toProductDomain(product ProductEntity) domain.Product {
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
	entities []ProductEntity,
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
