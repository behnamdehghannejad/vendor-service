package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
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

func (i *Inventory) ValidateID(idStr string) error {
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

func (i *Inventory) AddProductsToVendor(r dto.AddProductsToVendorRequest) error {
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
