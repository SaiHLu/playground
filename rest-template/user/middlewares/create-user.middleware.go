package middlewares

import (
	cm "rest-template/common/middlewares"
	"rest-template/user/dto"

	"github.com/gofiber/fiber/v2"
)

func CreateUserMiddleware(c *fiber.Ctx) error {
	body := new(dto.CreateUserDto)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors, ok := cm.ValidateRequest(body)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
