package auth

import (
	"rest-template/auth/dto"
	"rest-template/auth/services"

	"github.com/gofiber/fiber/v2"
)

func Login(service services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.LoginDto
		c.BodyParser(&requestBody)

		token, err := service.Login(requestBody)
		if err != nil {
			return c.JSON(dto.LoginErrorResponse(err))
		}

		return c.JSON(dto.LoginSuccessResponse(token))
	}
}
