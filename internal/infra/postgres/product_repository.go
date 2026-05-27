package postgres

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"

	"gorm.io/gorm"
)

type ProductEntity struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	Active      bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

func (repo *ProductRepository) Add(domain domain.Product) error {
	return repo.db.Save(repo.toProductEntity(domain)).Error
}

func (repo *ProductRepository) Update(domain domain.Product) error {
	return repo.db.Save(repo.toProductEntity(domain)).Error
}

func (repo *ProductRepository) Delete(id int) error {
	err := repo.db.Model(&ProductEntity{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"active":     false,
		}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err, id)
	}
	return nil
}

func (repo *ProductRepository) FindById(id int) (domain.Product, error) {
	var product ProductEntity
	if err := repo.db.Where("id = ?", id).First(&product).Error; err != nil {
		return domain.Product{}, convertPostgresErrorToAppError(err)
	}

	return repo.toProductDomain(product), nil
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
