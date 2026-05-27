package service

import (
	"time"
	"vendor-service/internal/domain"
)

func CreateHistoryByOrder(order domain.Order) *domain.History {
	return &domain.History{
		OrderID:   order.OrderID,
		Quantity:  order.Quantity,
		ProductID: order.ProductID,
		VendorID:  order.VendorID,
		Status:    domain.CREATED,
		Active:    true,
		CreatedAt: time.Now(),
	}
}
