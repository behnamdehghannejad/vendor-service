package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type ProductRepository interface {
	Add(domain.Product) error
	Update(domain.Product) error
	Delete(int) error
	FindById(int) (domain.Product, error)
}
