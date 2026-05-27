package validator

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var iranPhoneRegex = regexp.MustCompile(`^(?:\+98|98|0)?9\d{9}$`)

func validateIranPhone(value interface{}) error {
	s, _ := value.(string)

	if !iranPhoneRegex.MatchString(s) {
		return validation.NewError(
			"validation_phone",
			"invalid iranian phone number",
		)
	}

	return nil
}
