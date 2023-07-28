package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/playground/fiber/database"
	"github.com/playground/fiber/models"
	"github.com/playground/fiber/response"
)

func CreateUser(c *fiber.Ctx) error {

	user := new(models.User)
	log.Info("user: ", user)

	// log.Fatal("hello")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(response.Response{Message: err.Error(), Data: nil})
	}

	if err := database.DB.Db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(response.Response{Message: err.Error(), Data: nil})
	}

	return c.Status(200).JSON(response.Response{Message: "OK", Data: user})
}
