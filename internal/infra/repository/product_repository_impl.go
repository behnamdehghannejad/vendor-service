package repository

import (
	"time"
	"vendor-service/internal/domain"

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

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (repo *ProductRepositoryImpl) Add(domain *domain.Product) error {
	return repo.db.Save(toProductEntity(domain)).Error
}

func (repo *ProductRepositoryImpl) Update(domain *domain.Product) error {
	return repo.db.Save(toProductEntity(domain)).Error
}

func (repo *ProductRepositoryImpl) Delete(id int) error {
	var product ProductEntity
	if err := repo.db.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}

	product.UpdatedAt = time.Now()
	product.Active = false
	return repo.db.Save(&product).Error
}

func (repo *ProductRepositoryImpl) FindById(id int) (*domain.Product, error) {
	var product ProductEntity
	if err := repo.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return toProductDomain(&product), nil
}
