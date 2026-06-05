package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/adapter/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Category struct {
	category port.CategoryService
}

func NewCategory(categoryService port.CategoryService) *Category {
	return &Category{
		category: categoryService,
	}
}

func (v *Category) Delete(id int) error {
	_, err := v.category.FindById(id)
	return err
}

func (v *Category) Create(r dto.CreateCategoryRequest) error {
	err := validation.ValidateStruct(&r,

		validation.Field(&r.Name,
			validation.Required.Error("name is required"),
			validation.Length(2, 100).Error("name must be between 2 and 100 characters"),
		),
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

func (v *Category) Update(r dto.RequestUpdateCategory) error {
	err := validation.ValidateStruct(&r,

		validation.Field(&r.Name,
			validation.Required.Error("name is required"),
			validation.Length(2, 100).Error("name must be between 2 and 100 characters"),
		),
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
