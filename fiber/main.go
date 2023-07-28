package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/playground/fiber/database"
	"github.com/playground/fiber/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Could not load Env file.", err.Error())
	}
	database.SetupDB()

	app := fiber.New()

	routes.SetUpRoutes(app)

	app.Use(cors.New())

	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.SendStatus(500)
	// })

	app.Listen(":8080")
}
