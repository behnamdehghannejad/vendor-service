package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/behnamdehghannejad/vendor/internal/domain"
	"github.com/behnamdehghannejad/vendor/internal/handler/dto"
	"github.com/behnamdehghannejad/vendor/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) Create(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.CreateProductRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Create(toProductDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer)
}

func (handler *ProductHandler) GetById(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	product, err := handler.service.FindById(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toProductResponse(product))
}

func (handler *ProductHandler) Update(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.UpdateProductRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Update(toUpdateProductDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer)
}

func (handler *ProductHandler) Delete(writer http.ResponseWriter, request *http.Request) {
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

func toProductDomain(req dto.CreateProductRequest) *domain.Product {
	return &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Active:      true,
	}
}

func toUpdateProductDomain(req dto.UpdateProductRequest) *domain.Product {
	return &domain.Product{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Active:      req.Active,
	}
}

func toProductResponse(product *domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
