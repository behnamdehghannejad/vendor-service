package postgres

import (
	"errors"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func convertPostgresErrorToAppError(err error, inputs ...any) error {
	var pgErr *pgconn.PgError

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.Wrap(err).
			Input(inputs...).
			NotFound().
			Warning().
			Log().
			Build()
	}

	if errors.As(err, &pgErr) {
		return generatePostgresError(pgErr, inputs...)
	}

	return apperror.Wrap(err).
		Input(inputs...).
		UnExpected().
		Warning().
		Log().
		Build()
}

func generatePostgresError(pgErr *pgconn.PgError, inputs ...any) error {
	switch pgErr.Code {

	case "23505":
		return apperror.Wrap(pgErr).
			Input(inputs...).
			Duplicate().
			Warning().
			Log().
			Build()

	case "23503":
		return apperror.Wrap(pgErr).
			Input(inputs...).
			BadRequest().
			Warning().
			Log().
			Build()

	case "23502":
		return apperror.Wrap(pgErr).
			Input(inputs...).
			BadRequest().
			Warning().
			Log().
			Build()

	default:
		return apperror.Wrap(pgErr).
			Input(inputs...).
			UnExpected().
			Warning().
			Log().
			Build()
	}
}
