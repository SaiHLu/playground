package auth

import (
	"rest-template/auth/middlewares"
	"rest-template/auth/services"
	"rest-template/database"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) fiber.Router {
	api := app.Group("/auth")

	authRepository := services.NewAuthRepository(database.DB)
	authService := services.NewAuthService(authRepository)
	api.Post("/login", middlewares.LoginMiddleware, Login(authService))

	return api
}
