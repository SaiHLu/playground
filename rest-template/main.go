package main

import (
	"rest-template/auth"
	"rest-template/database"
	"rest-template/user"
	"rest-template/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())

	database.ConnectToDB()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "OK",
		})
	})

	app.Get("/rooms/:roomId", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	user.UsersRoute(app)
	auth.AuthRoute(app)
	ws.WsRoute(app)

	app.Listen(":8080")
}
