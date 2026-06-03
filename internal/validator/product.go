package validator

import (
	"strconv"

	"github.com/behnamdehghannejad/vendorservice/internal/adapter/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Product struct {
	productService port.ProductService
}

func NewProduct(productService port.ProductService) *Product {
	return &Product{
		productService: productService,
	}
}

func (p *Product) ValidateID(idStr string) error {
	err := validation.Validate(
		idStr,
		validation.Required,
		is.Digit,
	)
	if err != nil {
		return apperror.
			Wrap(err).
			BadRequest().
			Log().
			Build()
	}
	return nil
}

func (p *Product) Create(r dto.CreateProductRequest) error {
	err := validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(
			&r.Description,
			validation.Required,
			validation.Length(10, 1000),
		),
	)
	if err != nil {
		apperror.Wrap(err).
			BadRequest().
			Build()
	}
	return nil
}

func (p *Product) Update(r dto.RequestUpdateProduct, idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return apperror.Wrap(err).BadRequest().Build()
	}

	if _, err := p.productService.FindById(id); err != nil {
		return err
	}

	err = validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(
			&r.Description,
			validation.Required,
			validation.Length(10, 1000),
		),
	)
	if err != nil {
		apperror.Wrap(err).
			BadRequest().
			Build()
	}
	return nil
}
