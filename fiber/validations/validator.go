package validations

import "github.com/go-playground/validator/v10"

var Validator = validator.New()

type Error struct {
	Field string
	Tag   string
	Value string
}
