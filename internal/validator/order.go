package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Order struct {
	service port.OrderService
}

func NewOrder(service port.OrderService) *Order {
	return &Order{
		service: service,
	}
}

func (o *Order) ValidateID(idStr string) error {
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

func (h *Order) Create(r dto.ManageOrdersRequest) error {
	err := validation.ValidateStruct(&r)
	if err != nil {
		return apperror.
			Wrap(err).
			BadRequest().
			Log().
			Build()
	}

	return nil
}
