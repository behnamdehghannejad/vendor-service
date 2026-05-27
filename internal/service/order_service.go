package service

import "vendor-service/internal/domain"

type OrderService interface {
	ManageOrders(orders *domain.ListOrder) error
	AcceptOrdersPayment(orders *domain.ListOrder) error
}
