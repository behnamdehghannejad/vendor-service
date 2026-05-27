package httphandler

import (
	"net/http"
	"strconv"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	service port.VendorService
}

func NewVendorHandler(service port.VendorService) *VendorHandler {
	return &VendorHandler{service: service}
}

func (handler *VendorHandler) Create(c *gin.Context) {
	var req dto.CreateVendorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
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
		errorresponse, status := httperror.Handle(err)
		c.JSON(status, errorresponse)
		return
	}

	vendor, err := handler.service.FindByID(id)
	if err != nil {
		errresponse, status := httperror.Handle(err)
		c.JSON(status, errresponse)
		return
	}

	c.JSON(http.StatusOK, toVendorResponse(vendor))
}

func (handler *VendorHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorresponse, status := httperror.Handle(err)
		c.JSON(status, errorresponse)
		return
	}

	if err := handler.service.Delete(id); err != nil {
		errorresponse, status := httperror.Handle(err)
		c.JSON(status, errorresponse)
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
