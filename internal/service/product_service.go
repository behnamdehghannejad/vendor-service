package service

import "github.com/behnamdehghannejad/vendor/internal/domain"

type ProductService interface {
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int) error
	FindById(id int) (*domain.Product, error)
}
