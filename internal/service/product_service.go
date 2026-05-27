package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type ProductService struct {
	repository port.ProductRepository
}

func NewProductService(repository port.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (service *ProductService) Create(product domain.Product) error {
	return service.repository.Add(product)
}

func (service *ProductService) Update(product domain.Product) error {
	return service.repository.Update(product)
}

func (service *ProductService) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *ProductService) FindById(id int) (domain.Product, error) {
	return service.repository.FindById(id)
}
