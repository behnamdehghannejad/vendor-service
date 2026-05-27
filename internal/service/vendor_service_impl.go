package service

import (
	"github.com/behnamdehghannejad/vendor/internal/domain"

	"github.com/behnamdehghannejad/vendor/internal/infra/repository"
)

type VendorServiceImpl struct {
	repository repository.VendorRepository
}

func NewVendorService(repository repository.VendorRepository) *VendorServiceImpl {
	return &VendorServiceImpl{repository: repository}
}

func (service *VendorServiceImpl) Create(vendor *domain.Vendor) error {
	return service.repository.Add(vendor)
}

func (service *VendorServiceImpl) Update(vendor *domain.Vendor) error {
	return service.repository.Update(vendor)
}

func (service *VendorServiceImpl) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *VendorServiceImpl) FindByID(id int) (*domain.Vendor, error) {
	return service.repository.FindByID(id)
}

func (service *VendorServiceImpl) FindByCode(code string) (*domain.Vendor, error) {
	return service.repository.FindByCode(code)
}
