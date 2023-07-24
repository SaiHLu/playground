package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playground/h-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	users := types.User{
		// ID:        "123",
		FirstName: "Ko Sai",
		LastName:  "Hlaing Lu",
	}
	return c.JSON(users)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(c.Params("id"))
}
