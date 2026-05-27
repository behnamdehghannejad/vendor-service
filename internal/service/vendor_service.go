package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type VendorService struct {
	repository port.VendorRepository
}

func NewVendorService(repository port.VendorRepository) *VendorService {
	return &VendorService{repository: repository}
}

func (service *VendorService) Create(vendor domain.Vendor) error {
	return service.repository.Add(vendor)
}

func (service *VendorService) Update(vendor domain.Vendor) error {
	return service.repository.Update(vendor)
}

func (service *VendorService) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *VendorService) FindByID(id int) (domain.Vendor, error) {
	return service.repository.FindByID(id)
}

func (service *VendorService) FindByCode(code string) (domain.Vendor, error) {
	return service.repository.FindByCode(code)
}
