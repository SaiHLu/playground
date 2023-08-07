package dto

import "github.com/gofiber/fiber/v2"

type Auth struct {
	AccessToken string `json:"access_token"`
}

func LoginSuccessResponse(data *Auth) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    &data,
		"error":   nil,
	}
}

func LoginErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"data":    nil,
		"error":   err.Error(),
	}
}
