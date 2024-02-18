package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HttpHandleBindErrors(err error) []string {
	errors := make([]string, 0, 1)

	if _, ok := err.(validator.ValidationErrors); ok {
		for _, e := range err.(validator.ValidationErrors) {
			e := fmt.Sprintf("Field validation: %s, Error: %s", e.Field(), e.Tag())
			errors = append(errors, e)
		}

		return errors
	}

	errors = append(errors, err.Error())
	return errors
}
