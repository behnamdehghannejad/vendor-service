package validator

import (
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type History struct {
	history port.HistoryService
}

func NewHistory(history port.HistoryService) *History {
	return &History{
		history: history,
	}
}

func (h *History) ValidateID(idStr string) error {
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

func (h *History) Create(r dto.CreateHistoryRequest) error {
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
