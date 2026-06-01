package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type ProductRepository interface {
	Create(domain.Product) error
	Update(domain.Product) error
	Delete(int) error
	FindById(int) (domain.Product, error)
	Filter(domain.SearchProduct) ([]domain.Product, error)
}
