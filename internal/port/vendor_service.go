package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type VendorService interface {
	Create(domain.Vendor) error
	Update(domain.Vendor) error
	Delete(int) error
	FindByID(int) (domain.Vendor, error)
	Filter(domain.SearchVendor) ([]domain.Vendor, error)
	IsActive(id int) error
}
