package httphandler

import (
	"net/http"
	"strconv"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	service   port.VendorService
	validator *validator.Vendor
}

func NewVendorHandler(service port.VendorService, validator *validator.Vendor) *VendorHandler {
	return &VendorHandler{
		service:   service,
		validator: validator,
	}
}

func (handler *VendorHandler) Create(c *gin.Context) {
	var req dto.CreateVendorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := handler.validator.Create(req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	err := handler.service.Create(toVendorDomain(req))
	if err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusCreated)
}

func (handler *VendorHandler) GetById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	vendor, err := handler.service.FindByID(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, toVendorResponse(vendor))
}

func (handler *VendorHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := handler.validator.Delete(id); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := handler.service.Delete(id); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.Status(http.StatusOK)
}

func toVendorDomain(req dto.CreateVendorRequest) domain.Vendor {
	return domain.Vendor{
		Code:    req.Code,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Active:  true,
	}
}

func toVendorResponse(vendor domain.Vendor) dto.VendorResponse {
	return dto.VendorResponse{
		ID:        vendor.ID,
		Code:      vendor.Code,
		Name:      vendor.Name,
		Email:     vendor.Email,
		Phone:     vendor.Phone,
		Address:   vendor.Address,
		Active:    vendor.Active,
		CreatedAt: vendor.CreatedAt,
		UpdatedAt: vendor.UpdatedAt,
	}
}
