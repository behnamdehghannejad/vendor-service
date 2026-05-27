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

func (repo *VendorRepository) Add(v domain.Vendor) error {
	entity := repo.toVendorEntity(v)

	if err := repo.db.Create(&entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, v)
	}

	return nil
}

func (repo *VendorRepository) Update(v domain.Vendor) error {
	entity := repo.toVendorEntity(v)

	if err := repo.db.Save(&entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, v)
	}

	return nil
}

func (repo *VendorRepository) Delete(id int) error {
	if err := repo.db.Model(&VendorEntity{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"active":     false,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return convertPostgresErrorToAppError(err)
	}

	return nil
}

func (repo *VendorRepository) Filter(filter domain.SearchVendor) ([]domain.Vendor, error) {
	var entities []VendorEntity

	query := repo.db.Model(&VendorEntity{})

	if filter.IsActive != nil {
		query = query.Where("active = ?", *filter.IsActive)
	}

	if filter.Code != "" {
		query = query.Where("code = ?", filter.Code)
	}

	if filter.SearchName != "" {
		query = query.Where("name ILIKE ?", "%"+filter.SearchName+"%")
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	return repo.toVendorsDomain(entities), nil
}

func (repo *VendorRepository) FindByID(id int) (domain.Vendor, error) {
	var entity VendorEntity

	if err := repo.db.First(&entity, id).Error; err != nil {
		return domain.Vendor{}, convertPostgresErrorToAppError(err)
	}

	return repo.toVendorDomain(entity), nil
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

func (repo *VendorRepository) toVendorsDomain(vendors []VendorEntity) []domain.Vendor {
	vendorsDomain := make([]domain.Vendor, 0, len(vendors))
	for _, vendor := range vendors {
		vendorsDomain = append(vendorsDomain, repo.toVendorDomain(vendor))
	}
	return vendorsDomain
}
