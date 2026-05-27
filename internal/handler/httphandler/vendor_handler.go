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

func (h *VendorHandler) Create(c *gin.Context) {
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

	err := h.service.Create(h.toVendorDomain(req))
	if err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *VendorHandler) GetById(c *gin.Context) {
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

	c.JSON(http.StatusOK, h.toVendorResponse(vendor))
}

func (h *VendorHandler) Delete(c *gin.Context) {
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

	c.Status(http.StatusOK)
}

func (h *VendorHandler) Filter(c *gin.Context) {
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

	c.JSON(http.StatusOK, h.toVendorsResponse(vendors))
}

func (h *VendorHandler) GetIsActiveFromQuery(activeStr string) *bool {
	active := true
	deActive := false
	switch activeStr {
	case "active":
		return &active
	case "deActive":
		return &deActive
	}
	return nil
}

func (h *VendorHandler) toVendorDomain(req dto.CreateVendorRequest) domain.Vendor {
	return domain.Vendor{
		Code:    req.Code,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Active:  true,
	}
}

func (h *VendorHandler) toVendorResponse(vendor domain.Vendor) dto.VendorResponse {
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

func (h *VendorHandler) toVendorsResponse(vendors []domain.Vendor) dto.VendorsResponse {
	vendorsResponse := make([]dto.VendorResponse, 0, len(vendors))
	for _, vendor := range vendors {
		vendorsResponse = append(vendorsResponse, h.toVendorResponse(vendor))
	}

	return dto.VendorsResponse{
		Items: vendorsResponse,
	}
}
