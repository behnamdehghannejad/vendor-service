package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/adapter/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Inventory struct {
	inventory port.InventoryService
}

func NewInventory(inventory port.InventoryService) *Inventory {
	return &Inventory{
		inventory: inventory,
	}
}

func (i *Inventory) ValidateIDs(vendorID string, productID string) error {
	err := validation.Validate(vendorID,
		validation.Required,
		is.Digit,
	)
	if err != nil {
		return apperror.Wrap(err).
			BadRequest().
			Build()
	}

	err = validation.Validate(productID,
		validation.Required,
		is.Digit,
	)
	if err != nil {
		return apperror.Wrap(err).
			BadRequest().
			Build()
	}

	return nil
}

func (i *Inventory) Upsert(req dto.RequestUpsertInventory) error {
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Quantity,
			validation.Required,
			validation.Min(0),
		),
	)
	if err != nil {
		return apperror.Wrap(err).
			BadRequest().
			Build()
	}
	return nil
}
