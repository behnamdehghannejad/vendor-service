package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type OrderService interface {
	ManageOrders(orders domain.ListOrder) error
	AcceptOrdersPayment(orders domain.ListOrder) error
}
