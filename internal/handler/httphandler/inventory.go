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

type Inventory struct {
	service   port.InventoryService
	validator *validator.Inventory
}

func NewInventoryHandler(service port.InventoryService, validator *validator.Inventory) *Inventory {
	return &Inventory{
		service:   service,
		validator: validator,
	}
}

func (h *Inventory) AddProductsToVendor(c *gin.Context) {
	var req dto.AddProductsToVendorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := h.validator.AddProductsToVendor(req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	err := h.service.AddProductsToVendor(h.toInventoryDomain(req))
	if err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Inventory) toInventoryDomain(req dto.AddProductsToVendorRequest) domain.Inventory {
	return domain.Inventory{
		VendorID:  req.VendorID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Reserved:  0,
	}
}
