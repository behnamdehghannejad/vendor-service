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

type Product struct {
	productService port.ProductService
	validator      *validator.Product
}

func NewProductHandler(productSvc port.ProductService, validator *validator.Product) *Product {
	return &Product{
		productService: productSvc,
	}
}

func (h *Product) Create(c *gin.Context) {
	var req dto.CreateProductRequest

	if err := c.ShouldBind(&req); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	err := h.productService.Create(domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Active:      true,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
	}

	c.Status(http.StatusCreated)
}

func (h *Product) GetById(c *gin.Context) {
	idStr := c.Query("id")

	err := h.validator.FindById(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	id, _ := strconv.Atoi(idStr)
	product, err := h.productService.FindById(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}
	c.JSON(http.StatusOK, h.serializeProduct(product))
}

func (h *Product) Filter(c *gin.Context) {
	searchName := c.Query("search_name")

	products, err := h.productService.Filter(domain.SearchProduct{
		SearchName: searchName,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeProducts(products))
}

func (h *Product) serializeProducts(products []domain.Product) []dto.ProductResponse {
	productsResponse := make([]dto.ProductResponse, 0, len(products))

	for _, product := range products {
		productsResponse = append(productsResponse, h.serializeProduct(product))
	}
	return productsResponse
}

func (h *Product) serializeProduct(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
