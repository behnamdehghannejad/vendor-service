package service

import (
	"vendor-service/internal/domain"
	"vendor-service/internal/infra/repository"
)

type OrderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{repository: repository}
}

func (service *OrderServiceImpl) Create(order domain.Order) error {
	return service.repository.Create(order)
}

func (service *OrderServiceImpl) GetByID(id int) (order domain.Order, err error) {
	order, err = service.repository.GetById(id)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (service *OrderServiceImpl) GetByUserId(userId int) (order domain.Order, err error) {
	order, err = service.repository.GetByUserId(userId)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (service *OrderServiceImpl) UpdateStatus(id int, status domain.Status) (order domain.Order, err error) {
	order, err = service.repository.UpdateStatus(id, status)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (service *OrderServiceImpl) ListAll() (orders []domain.Order, err error) {
	orders, err = service.repository.AllOrders()
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}

func (service *OrderServiceImpl) Delete(id int) error {
	if err := service.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
