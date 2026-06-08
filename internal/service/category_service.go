package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type CategoryService struct {
	repository port.CategoryRepository
}

func NewCategoryService(
	repository port.CategoryRepository,
) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) Create(category domain.Category) (int, error) {
	return s.repository.Create(category)
}

func (s *CategoryService) Update(category domain.Category) error {
	return s.repository.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repository.SoftDelete(id)
}

func (s *CategoryService) FindById(id int) (domain.Category, error) {
	return s.repository.FindById(id)
}

func (s *CategoryService) FindChildren(id int) ([]domain.Category, error) {
	return s.repository.FindChildren(id)
}

func (s *CategoryService) FindParents(id int) ([]domain.Category, error) {
	return s.repository.FindParents(id)
}

func (s *CategoryService) IsActive(id int) error {
	category, err := s.FindById(id)
	if err != nil {
		return err
	}

	return domain.IsActiveCategory(category.Active)
}
