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

func (s *VendorService) Create(vendor domain.Vendor) error {
	return s.repository.Create(vendor)
}

func (s *VendorService) Update(vendor domain.Vendor) error {
	return s.repository.Update(vendor)
}

func (s *VendorService) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *VendorService) FindByID(id int) (domain.Vendor, error) {
	return s.repository.FindByID(id)
}

func (s *VendorService) Filter(filter domain.SearchVendor) ([]domain.Vendor, error) {
	return s.repository.Filter(filter)
}

func (s *VendorService) IsActive(id int) error {
	vendor, err := s.FindByID(id)
	if err != nil {
		return err
	}

	return domain.IsActiveVendor(vendor.Active)
}
