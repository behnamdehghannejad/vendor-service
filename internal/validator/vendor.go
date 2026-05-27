package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Vendor struct {
	vendor port.VendorService
}

func NewVendorValidator(vendorService port.VendorService) *Vendor {
	return &Vendor{
		vendor: vendorService,
	}
}

func (v *Vendor) Delete(id int) error {
	_, err := v.vendor.FindByID(id)
	return err
}

func (v *Vendor) Create(r dto.CreateVendorRequest) error {
	err := validation.ValidateStruct(&r,

		validation.Field(&r.Code,
			validation.Required.Error("code is required"),
			validation.Length(3, 30).Error("code must be between 3 and 30 characters"),
		),

		validation.Field(&r.Name,
			validation.Required.Error("name is required"),
			validation.Length(2, 100).Error("name must be between 2 and 100 characters"),
		),

		validation.Field(&r.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),

		validation.Field(&r.Phone,
			validation.Required.Error("phone is required"),
			validation.By(validateIranPhone),
		),

		validation.Field(&r.Address,
			validation.Required.Error("address is required"),
			validation.Length(5, 200).Error("address must be between 5 and 200 characters"),
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
