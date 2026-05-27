package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/behnamdehghannejad/vendor/internal/domain"
	"github.com/behnamdehghannejad/vendor/internal/handler/dto"
	"github.com/behnamdehghannejad/vendor/internal/service"
)

type VendorHandler struct {
	service service.VendorService
}

func NewVendorHandler(service service.VendorService) *VendorHandler {
	return &VendorHandler{service: service}
}

func (handler *VendorHandler) Create(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.CreateVendorRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Create(toVendorDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer)
}

func (handler *VendorHandler) GetById(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	vendor, err := handler.service.FindByID(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	returnVendorResponse(writer, vendor)
}

func (handler *VendorHandler) GetByCode(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	code := request.PathValue("code")

	vendor, err := handler.service.FindByCode(code)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	returnVendorResponse(writer, vendor)
}

func (handler *VendorHandler) Update(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.UpdateVendorRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Update(toUpdateVendorDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer)
}

func (handler *VendorHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	if err := handler.service.Delete(id); err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer)
}

func toVendorDomain(req dto.CreateVendorRequest) *domain.Vendor {
	return &domain.Vendor{
		Code:    req.Code,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Active:  true,
	}
}

func toUpdateVendorDomain(req dto.UpdateVendorRequest) *domain.Vendor {
	return &domain.Vendor{
		ID:      req.ID,
		Code:    req.Code,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Active:  req.Active,
	}
}

func returnVendorResponse(writer http.ResponseWriter, vendor *domain.Vendor) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toVendorResponse(vendor))
}

func toVendorResponse(vendor *domain.Vendor) dto.VendorResponse {
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
