package validator

import (
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

func (i *Inventory) ValidateIDs(ID string) error {
	err := validation.Validate(ID,
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
