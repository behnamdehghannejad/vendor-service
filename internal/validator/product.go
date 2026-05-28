package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Product struct{}

func NewProduct() *Product {
	return &Product{}
}

func (v *Product) ValidateID(idStr string) error {
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
