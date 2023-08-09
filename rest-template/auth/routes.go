package auth

import (
	"rest-template/auth/controllers"
	"rest-template/auth/middlewares"
	"rest-template/auth/repositories"
	"rest-template/auth/services"
	"rest-template/database"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) fiber.Router {
	api := app.Group("/auth")

	authRepository := repositories.NewAuthRepository(database.DB)
	authService := services.NewAuthService(authRepository)
	api.Post("/login", middlewares.LoginMiddleware, controllers.Login(authService))

	return api
}
