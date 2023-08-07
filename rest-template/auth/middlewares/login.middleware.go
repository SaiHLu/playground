package middlewares

import (
	"rest-template/auth/dto"
	"rest-template/common/middlewares"

	"github.com/gofiber/fiber/v2"
)

func LoginMiddleware(c *fiber.Ctx) error {
	body := new(dto.LoginDto)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors, ok := middlewares.ValidateRequest(body)

	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
