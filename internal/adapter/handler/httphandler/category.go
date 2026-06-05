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

type Category struct {
	service   port.CategoryService
	validator *validator.Category
}

func NewCategoryHandler(service port.CategoryService, validator *validator.Category) *Category {
	return &Category{
		service:   service,
		validator: validator,
	}
}

func (h *Category) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest

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

	_, err := h.service.Create(toCategoryDomain(req))
	if err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Category) Update(c *gin.Context) {
	var req dto.RequestUpdateCategory

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := h.validator.Update(req); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	if err := h.service.Update(toCategoryUpdateDomain(req)); err != nil {
		responseError, status := httperror.Handle(err)
		c.JSON(status, responseError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Category) GetById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	category, err := h.service.FindById(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, serializeCategory(category))
}

func (h *Category) Delete(c *gin.Context) {
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

func (h *Category) FindChildren(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	categories, err := h.service.FindChildren(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, serializeCategories(categories))
}

func (h *Category) FindParents(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	categories, err := h.service.FindParents(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, serializeCategories(categories))
}

func (h *Category) GetIsActiveFromQuery(activeStr string) *bool {
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

func toCategoryDomain(req dto.CreateCategoryRequest) domain.Category {
	return domain.Category{
		Name:   req.Name,
		Active: true,
	}
}

func toCategoryUpdateDomain(req dto.RequestUpdateCategory) domain.Category {
	return domain.Category{
		ID:   req.ID,
		Name: req.Name,
	}
}

func serializeCategory(category domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		ParentID: category.ParentID,
		Path:     category.Path,
	}
}

func serializeCategories(categories []domain.Category) []dto.CategoryResponse {
	var categoriesResponse []dto.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, serializeCategory(category))
	}
	return categoriesResponse
}
