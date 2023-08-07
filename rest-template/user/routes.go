package user

import (
	cm "rest-template/common/middlewares"
	"rest-template/database"
	"rest-template/user/middlewares"
	"rest-template/user/services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App) fiber.Router {
	userRepository := services.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)

	api := app.Group("/users", cm.Protected())

	api.Get("/", Find(userService))
	api.Post("/", middlewares.CreateUserMiddleware, Create(userService))
	api.Patch("/:id", middlewares.UpdateUserMiddleware, Update(userService))
	api.Delete("/:id", Delete(userService))

	return api
}
