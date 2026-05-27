package http

import (
	"encoding/json"
	"net/http"
	"vendor-service/internal/domain"
	"vendor-service/internal/handler/dto"
	"vendor-service/internal/service"
)

type InventoryHandler struct {
	service service.InventoryService
}

func NewInventoryHandler(service service.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

func (handler *InventoryHandler) AddProductsToVendor(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.AddProductsToVendorRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.AddProductsToVendor(toInventoryDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer)
}

func toInventoryDomain(req dto.AddProductsToVendorRequest) *domain.Inventory {
	return &domain.Inventory{
		Vendor: domain.Vendor{
			ID: req.VendorID,
		},
		Product: domain.Product{
			ID: req.ProductID,
		},
		Quantity: req.Quantity,
		Reserved: 0,
	}
}
