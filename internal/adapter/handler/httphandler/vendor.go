package httphandler

import (
	"net/http"
	"strconv"

	"github.com/behnamdehghannejad/vendorservice/internal/adapter/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
)

type Vendor struct {
	service   port.VendorService
	validator *validator.Vendor
}

func NewVendorHandler(service port.VendorService, validator *validator.Vendor) *Vendor {
	return &Vendor{
		service:   service,
		validator: validator,
	}
}

func (h *Vendor) Create(c *gin.Context) {
	var req dto.CreateVendorRequest

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

	_, err := h.service.Create(h.toVendorDomain(req))
	if err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Vendor) GetById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	vendor, err := h.service.FindByID(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeVendor(vendor))
}

func (h *Vendor) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := h.validator.Delete(id); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := h.service.Delete(id); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Vendor) Filter(c *gin.Context) {
	code := c.Query("code")
	searchName := c.Query("search_name")
	isActive := h.GetIsActiveFromQuery(c.Query("active"))

	vendors, err := h.service.Filter(domain.SearchVendor{
		Code:       code,
		SearchName: searchName,
		IsActive:   isActive,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeVendors(vendors))
}

func (h *Vendor) GetIsActiveFromQuery(activeStr string) *bool {
	active := true
	deActive := false
	switch activeStr {
	case "active":
		return &active
	case "deactive":
		return &deActive
	}
	return nil
}

func (h *Vendor) toVendorDomain(req dto.CreateVendorRequest) domain.Vendor {
	return domain.Vendor{
		Code:    req.Code,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Active:  true,
	}
}

func (h *Vendor) serializeVendor(vendor domain.Vendor) dto.VendorResponse {
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

func (h *Vendor) serializeVendors(vendors []domain.Vendor) dto.VendorsResponse {
	vendorsResponse := make([]dto.VendorResponse, 0, len(vendors))
	for _, vendor := range vendors {
		vendorsResponse = append(vendorsResponse, h.serializeVendor(vendor))
	}

	return dto.VendorsResponse{
		Items: vendorsResponse,
	}
}
