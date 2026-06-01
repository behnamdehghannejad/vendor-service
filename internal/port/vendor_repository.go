package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type VendorRepository interface {
	Create(domain.Vendor) error
	Update(domain.Vendor) error
	Delete(int) error
	Filter(domain.SearchVendor) ([]domain.Vendor, error)
	FindByID(int) (domain.Vendor, error)
}
