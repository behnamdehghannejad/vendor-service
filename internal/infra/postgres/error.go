package postgres

import (
	"errors"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func convertPostgresErrorToAppError(err error, inputs ...any) error {
	var pqErr *pq.Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.Wrap(err).
			Input(inputs...).
			NotFound().
			Warning().
			Log().
			Build()
	}

	if errors.As(err, &pqErr) {
		return generatePostgresError(pqErr, inputs...)
	}

	return apperror.Wrap(err).
		Input(inputs...).
		UnExpected().
		Warning().
		Log().
		Build()
}

func generatePostgresError(pqErr *pq.Error, inputs ...any) error {
	switch pqErr.Code {

	case "23505":
		return apperror.Wrap(pqErr).
			Input(inputs...).
			Duplicate().
			Warning().
			Log().
			Build()

	case "23503":
		return apperror.Wrap(pqErr).
			Input(inputs...).
			BadRequest().
			Warning().
			Log().
			Build()

	case "23502":
		return apperror.Wrap(pqErr).
			Input(inputs...).
			BadRequest().
			Warning().
			Log().
			Build()

	default:
		return apperror.Wrap(pqErr).
			Input(inputs...).
			UnExpected().
			Warning().
			Log().
			Build()
	}
}
