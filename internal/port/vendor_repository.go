package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type VendorRepository interface {
	Create(domain.Vendor) (int, error)
	DeleteVendorsByIDs(...int) error
	Update(domain.Vendor) error
	SoftDelete(int) error
	Filter(domain.SearchVendor) ([]domain.Vendor, error)
	FindByID(int) (domain.Vendor, error)
}
