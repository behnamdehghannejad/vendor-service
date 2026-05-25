package service

import (
	"vendor-service/internal/domain"
	"vendor-service/internal/infra/repository"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{repository: repository}
}

func (service *ProductServiceImpl) Create(product *domain.Product) error {
	return service.repository.Add(product)
}

func (service *ProductServiceImpl) Update(product *domain.Product) error {
	return service.repository.Update(product)
}

func (service *ProductServiceImpl) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *ProductServiceImpl) FindById(id int) (*domain.Product, error) {
	return service.repository.FindById(id)
}
