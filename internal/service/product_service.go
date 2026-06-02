package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type ProductService struct {
	repository port.ProductRepository
}

func NewProductService(
	repository port.ProductRepository,
) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Create(product domain.Product) (int, error) {
	return s.repository.Create(product)
}

func (s *ProductService) Update(product domain.Product) error {
	return s.repository.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repository.SoftDelete(id)
}

func (s *ProductService) FindById(id int) (domain.Product, error) {
	return s.repository.FindById(id)
}

func (s *ProductService) Filter(filter domain.SearchProduct) ([]domain.Product, error) {
	return s.repository.Filter(filter)
}

func (s *ProductService) IsActive(id int) error {
	product, err := s.FindById(id)
	if err != nil {
		return err
	}

	return domain.IsActiveProduct(product.Active)
}

func (*ProductService) getProductIDs(products []domain.Product) []int {
	IDs := make([]int, 0, len(products))
	for _, product := range products {
		IDs = append(IDs, product.ID)
	}
	return IDs
}
