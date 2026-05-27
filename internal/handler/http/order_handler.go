package http

import (
	"encoding/json"
	"net/http"

	"vendor-service/internal/domain"
	"vendor-service/internal/handler/dto"
	"vendor-service/internal/service"

	"github.com/google/uuid"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (handler *OrderHandler) AddOrders(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.ManageOrdersRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.service.ManageOrders(toListOrderDomain(req)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer)
}

func (handler *OrderHandler) AcceptOrdersPayment(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.AcceptOrdersPaymentRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.service.AcceptOrdersPayment(toAcceptOrdersDomain(req)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer)
}

func toListOrderDomain(req dto.ManageOrdersRequest) *domain.ListOrder {
	orders := make([]domain.Order, 0, len(req.Orders))

	for _, order := range req.Orders {
		orderID, _ := uuid.Parse(order.OrderID)
		orders = append(orders, domain.Order{
			OrderID:   orderID,
			VendorID:  order.VendorID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
		})
	}

	return &domain.ListOrder{
		Orders: orders,
	}
}

func toAcceptOrdersDomain(req dto.AcceptOrdersPaymentRequest) *domain.ListOrder {
	orders := make([]domain.Order, 0, len(req.Orders))

	for _, order := range req.Orders {
		orderID, _ := uuid.Parse(order.OrderID)
		paymentID, _ := uuid.Parse(order.OrderID)
		orders = append(
			orders,
			domain.Order{
				OrderID:   orderID,
				PaymentID: paymentID,
				VendorID:  order.VendorID,
				ProductID: order.ProductID,
				Quantity:  order.Quantity,
			},
		)
	}

	return &domain.ListOrder{
		Orders: orders,
	}
}
