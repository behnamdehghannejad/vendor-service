package service

import (
	"vendor-service/internal/domain"
)

type OrderService interface {
	Create(order domain.Order) error
	GetByID(id int) (domain.Order, error)
	UpdateStatus(id int, status domain.Status) (domain.Order, error)
	GetByUserId(id int) (domain.Order, error)
	ListAll() ([]domain.Order, error)
	Delete(id int) error
}
