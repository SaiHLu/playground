package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/playground/fiber/models"
)

func CreateUser(c *fiber.Ctx) error {
	var errors []*Error

	body := new(models.User)
	c.BodyParser(&body)

	fmt.Println("body: ", body)

	err := Validator.Struct(body)
	if err != nil {
		fmt.Printf("Errors: %+v\n", err.(validator.ValidationErrors))
		for _, err := range err.(validator.ValidationErrors) {
			var el Error
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}

		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
