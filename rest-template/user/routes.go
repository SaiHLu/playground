package user

import (
	cm "rest-template/common/middlewares"
	"rest-template/database"
	"rest-template/user/controllers"
	"rest-template/user/middlewares"
	"rest-template/user/repositories"
	"rest-template/user/services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App) fiber.Router {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)

	api := app.Group("/users", cm.Protected())

	api.Get("/", controllers.Find(userService))
	api.Post("/", middlewares.CreateUserMiddleware, controllers.Create(userService))
	api.Patch("/:id", middlewares.UpdateUserMiddleware, controllers.Update(userService))
	api.Delete("/:id", controllers.Delete(userService))

	return api
}
