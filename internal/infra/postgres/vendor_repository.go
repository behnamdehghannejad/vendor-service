package postgres

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"

	"gorm.io/gorm"
)

type VendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{
		db: db,
	}
}

func (repo *VendorRepository) Create(v domain.Vendor) (int, error) {
	vendorModel := repo.toVendorModel(v)

	if err := repo.db.Create(&vendorModel).Error; err != nil {
		return 0, convertPostgresErrorToAppError(err, v)
	}

	return vendorModel.ID, nil
}

func (repo *VendorRepository) Update(v domain.Vendor) error {
	entity := repo.toVendorModel(v)

	if err := repo.db.Save(&entity).Error; err != nil {
		return convertPostgresErrorToAppError(err, v)
	}

	return nil
}

func (repo *VendorRepository) SoftDelete(id int) error {
	if err := repo.db.Model(&model.VendorModel{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"active":     false,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return convertPostgresErrorToAppError(err)
	}

	return nil
}

func (repo *VendorRepository) DeleteVendorsByIDs(IDs ...int) error {
	err := repo.db.
		Where("id IN ?", IDs).
		Delete(&model.VendorModel{}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *VendorRepository) Filter(filter domain.SearchVendor) ([]domain.Vendor, error) {
	var entities []model.VendorModel

	query := repo.db.Model(&model.VendorModel{})

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
	var entity model.VendorModel

	if err := repo.db.First(&entity, id).Error; err != nil {
		return domain.Vendor{}, convertPostgresErrorToAppError(err)
	}

	return repo.toVendorDomain(entity), nil
}

func (repo *VendorRepository) toVendorModel(vendor domain.Vendor) model.VendorModel {
	return model.VendorModel{
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

func (repo *VendorRepository) toVendorDomain(vendor model.VendorModel) domain.Vendor {
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

func (repo *VendorRepository) toVendorsDomain(vendors []model.VendorModel) []domain.Vendor {
	vendorsDomain := make([]domain.Vendor, 0, len(vendors))
	for _, vendor := range vendors {
		vendorsDomain = append(vendorsDomain, repo.toVendorDomain(vendor))
	}
	return vendorsDomain
}
