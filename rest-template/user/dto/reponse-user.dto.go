package dto

import (
	"rest-template/user/models"

	"github.com/gofiber/fiber/v2"
)

func UserSuccessReponse(data *models.User) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
		"error":   nil,
	}
}

func UsersSuccessResponse(data *[]models.User) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
		"error":   nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"data":    nil,
		"error":   err.Error(),
	}
}
