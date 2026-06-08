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

	_, err := h.productService.Create(domain.Product{
		Name:        req.Name,
		CategoryID:  req.CategoryID,
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
	idStr := c.Param("id")

	err := h.validator.ValidateID(idStr)
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

func (h *Product) GetProductByCategory(c *gin.Context) {
	idStr := c.Param("category_id")

	err := h.validator.ValidateID(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	id, _ := strconv.Atoi(idStr)
	product, err := h.productService.FindByCategoryId(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}
	c.JSON(http.StatusOK, h.serializeProducts(product))
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

func (h *Product) Update(c *gin.Context) {
	idStr := c.Param("id")

	var req dto.RequestUpdateProduct
	if err := c.ShouldBind(&req); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := h.validator.Update(req, idStr); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := h.validator.ValidateID(idStr); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	h.productService.Update(domain.Product{})
}

func (h *Product) serializeProducts(products []domain.Product) dto.ProductsResponse {
	productsResponse := make([]dto.ProductResponse, 0, len(products))

	for _, product := range products {
		productsResponse = append(productsResponse, h.serializeProduct(product))
	}

	return dto.ProductsResponse{
		Items: productsResponse,
	}
}

func (h *Product) serializeProduct(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
