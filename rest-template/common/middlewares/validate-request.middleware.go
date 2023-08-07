package middlewares

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest(i interface{}) (map[string]string, bool) {
	errors := make(map[string]string)
	v := validator.New()

	if err := v.Struct(i); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "email" {
				errors[strings.ToLower(err.Field())] = "Invalid E-mail format."
				continue
			}

			errors[strings.ToLower(err.Field())] = fmt.Sprintf("%s is %s %s", err.Field(), err.Tag(), err.Param())
		}

		return errors, false
	}

	return nil, true
}
