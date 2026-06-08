package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type CategoryRepository interface {
	Create(domain.Category) (int, error)
	Update(domain.Category) error
	SoftDelete(int) error
	FindById(int) (domain.Category, error)
	FindChildren(int) ([]domain.Category, error)
	FindParents(int) ([]domain.Category, error)
	Delete(id int) error
}
