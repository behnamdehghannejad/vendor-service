package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type ProductRepository interface {
	Create(domain.Product) (int, error)
	Update(domain.Product) error
	SoftDelete(int) error
	FindById(int) (domain.Product, error)
	Filter(domain.SearchProduct) ([]domain.Product, error)
	UpdateProductDiscountPercentages([]domain.ProductDiscountPercentage) error
}
