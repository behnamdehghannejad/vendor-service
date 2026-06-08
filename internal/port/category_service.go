package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type CategoryService interface {
	Create(domain.Category) (int, error)
	Update(domain.Category) error
	Delete(int) error
	FindById(int) (domain.Category, error)
	FindChildren(int) ([]domain.Category, error)
	FindParents(int) ([]domain.Category, error)
}
