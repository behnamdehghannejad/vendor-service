package repository

import (
	"time"

	"github.com/behnamdehghannejad/vendor/internal/domain"

	"gorm.io/gorm"
)

type VendorEntity struct {
	ID        int       `gorm:"primaryKey"`
	Code      string    `gorm:"size:50"`
	Name      string    `gorm:"size:200"`
	Email     string    `gorm:"size:100"`
	Phone     string    `gorm:"size:20"`
	Address   string    `gorm:"size:500"`
	Active    bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:column"created_at"`
	UpdatedAt time.Time `gorm:column"updated_at"`
}

func (VendorEntity) TableName() string {
	return "vendors"
}

type VendorRepositoryImpl struct {
	db *gorm.DB
}

func NewVendorRepositoryImpl(db *gorm.DB) *VendorRepositoryImpl {
	return &VendorRepositoryImpl{
		db: db,
	}
}

func (repo *VendorRepositoryImpl) Add(domain *domain.Vendor) error {
	return repo.db.Create(toVendorEntity(domain)).Error
}

func (repo *VendorRepositoryImpl) Update(domain *domain.Vendor) error {
	return repo.db.Save(domain).Error
}

func (repo *VendorRepositoryImpl) Delete(id int) error {
	var vendor VendorEntity
	if err := repo.db.Where("id = ?", id).First(&vendor).Error; err != nil {
		return err
	}
	vendor.UpdatedAt = time.Now()
	vendor.Active = false
	return repo.db.Save(vendor).Error
}

func (repo *VendorRepositoryImpl) FindByID(id int) (*domain.Vendor, error) {
	var vendor VendorEntity
	if err := repo.db.Where("id = ?", id).First(&vendor).Error; err != nil {
		return nil, err
	}
	return toVendorDomain(&vendor), nil
}

func (repo *VendorRepositoryImpl) FindByCode(code string) (*domain.Vendor, error) {
	var vendor VendorEntity
	if err := repo.db.Where("code = ?", code).Find(&vendor).Error; err != nil {
		return nil, err
	}
	return toVendorDomain(&vendor), nil
}
