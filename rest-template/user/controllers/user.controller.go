package controllers

import (
	"fmt"
	"rest-template/user/dto"
	"rest-template/user/models"
	"rest-template/user/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Create(service services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.CreateUserDto
		c.BodyParser(&requestBody)

		result, err := service.Create(requestBody)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				return c.JSON(&fiber.Map{
					"success": false,
					"data":    nil,
					"error":   "User already exists.",
				})
			}
			return c.JSON(dto.UserErrorResponse(err))
		}

		return c.JSON(dto.UserSuccessReponse(result))
	}
}

func Find(service services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.Find()

		user := c.Locals("authUser").(*models.User)
		fmt.Println("User: ", user.Email)

		if err != nil {
			return c.JSON(dto.UserErrorResponse(err))
		}

		return c.JSON(dto.UsersSuccessResponse(result))
	}
}

func Update(service services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.UpdateUserDto
		c.BodyParser(&requestBody)

		result, err := service.Update(c.Params("id"), requestBody)

		if err != nil {
			return c.JSON(dto.UserErrorResponse(err))
		}

		return c.JSON(dto.UserSuccessReponse(result))
	}
}

func Delete(service services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := service.Delete(c.Params("id")); err != nil {
			return c.JSON(dto.UserErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    "Delete Success",
			"error":   nil,
		})
	}
}
