package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type ProductRepository interface {
	Add(product domain.Product) error
	Update(product domain.Product) error
	Delete(id int) error
	FindById(id int) (domain.Product, error)
}
