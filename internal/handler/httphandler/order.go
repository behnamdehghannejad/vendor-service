package httphandler

import (
	"net/http"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
)

type Order struct {
	service   port.OrderService
	validator *validator.Order
}

func NewOrderHandler(service port.OrderService, validator *validator.Order) *Order {
	return &Order{
		service:   service,
		validator: validator,
	}
}

func (h *Order) AddOrders(c *gin.Context) {
	var req dto.ManageOrdersRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := h.validator.Create(req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := h.service.ManageOrders(h.toListOrderDomain(req)); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Order) AcceptOrdersPayment(c *gin.Context) {
	var req dto.AcceptOrdersPaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := h.service.AcceptOrdersPayment(h.toAcceptOrdersDomain(req)); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Order) toListOrderDomain(req dto.ManageOrdersRequest) domain.ListOrder {
	orders := make([]domain.Order, 0, len(req.Orders))

	for _, order := range req.Orders {
		orders = append(orders, domain.Order{
			OrderID:   order.OrderID,
			VendorID:  order.VendorID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
		})
	}

	return domain.ListOrder{
		Orders: orders,
	}
}

func (h *Order) toAcceptOrdersDomain(req dto.AcceptOrdersPaymentRequest) domain.ListOrder {
	orders := make([]domain.Order, 0, len(req.Orders))

	for _, order := range req.Orders {
		orders = append(orders, domain.Order{
			OrderID:   order.OrderID,
			PaymentID: order.PaymentID,
			VendorID:  order.VendorID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
		})
	}

	return domain.ListOrder{
		Orders: orders,
	}
}
