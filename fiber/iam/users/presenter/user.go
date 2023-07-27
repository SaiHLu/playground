package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/playground/fiber/iam/users/models"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func UserSuccessResponse(data *models.User) *fiber.Map {
	user := &User{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
	}

	return &fiber.Map{
		"success": true,
		"data":    user,
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
		"data":    "",
		"error":   err.Error(),
	}
}
