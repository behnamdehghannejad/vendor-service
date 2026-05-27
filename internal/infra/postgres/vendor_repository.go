package postgres

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"

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

type VendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{
		db: db,
	}
}

func (repo *VendorRepository) Add(domain domain.Vendor) error {
	if err := repo.db.Create(repo.toVendorEntity(domain)).Error; err != nil {
		return convertPostgresErrorToAppError(err, domain)
	}
	return nil
}

func (repo *VendorRepository) Update(domain domain.Vendor) error {
	if err := repo.db.Save(domain).Error; err != nil {
		return convertPostgresErrorToAppError(err, domain)
	}
	return nil
}

func (repo *VendorRepository) Delete(id int) error {
	var vendor VendorEntity
	if err := repo.db.Where("id = ?", id).First(&vendor).Error; err != nil {
		return convertPostgresErrorToAppError(err, id)
	}
	vendor.UpdatedAt = time.Now()
	vendor.Active = false
	if err := repo.db.Delete(vendor).Error; err != nil {
		return convertPostgresErrorToAppError(err, id)
	}
	return nil
}

func (repo *VendorRepository) FindByID(id int) (domain.Vendor, error) {
	var vendor VendorEntity
	if err := repo.db.Where("id = ?", id).First(&vendor).Error; err != nil {
		return domain.Vendor{}, convertPostgresErrorToAppError(err)
	}
	return repo.toVendorDomain(vendor), nil
}

func (repo *VendorRepository) FindByCode(code string) (domain.Vendor, error) {
	var vendor VendorEntity
	if err := repo.db.Where("code = ?", code).Find(&vendor).Error; err != nil {
		return domain.Vendor{}, convertPostgresErrorToAppError(err)
	}
	return repo.toVendorDomain(vendor), nil
}

func (repo *VendorRepository) toVendorEntity(vendor domain.Vendor) VendorEntity {
	return VendorEntity{
		ID:        vendor.ID,
		Name:      vendor.Name,
		Code:      vendor.Code,
		Email:     vendor.Email,
		Phone:     vendor.Phone,
		Address:   vendor.Address,
		Active:    vendor.Active,
		CreatedAt: vendor.CreatedAt,
		UpdatedAt: vendor.UpdatedAt,
	}
}

func (repo *VendorRepository) toVendorDomain(vendor VendorEntity) domain.Vendor {
	return domain.Vendor{
		ID:        vendor.ID,
		Name:      vendor.Name,
		Code:      vendor.Code,
		Email:     vendor.Email,
		Phone:     vendor.Phone,
		Address:   vendor.Address,
		Active:    vendor.Active,
		CreatedAt: vendor.CreatedAt,
		UpdatedAt: vendor.UpdatedAt,
	}
}
