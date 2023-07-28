package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playground/fiber/handlers"
	"github.com/playground/fiber/validations"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/users", validations.CreateUser, handlers.CreateUser)
}
