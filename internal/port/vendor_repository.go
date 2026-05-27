package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type VendorRepository interface {
	Add(vendor domain.Vendor) error
	Update(vendor domain.Vendor) error
	Delete(id int) error
	FindByID(id int) (domain.Vendor, error)
	FindByCode(code string) (domain.Vendor, error)
}
