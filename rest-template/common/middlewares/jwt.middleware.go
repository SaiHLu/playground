package middlewares

import (
	"fmt"
	"rest-template/common/config"
	"rest-template/database"
	"rest-template/user/repositories"
	"rest-template/user/services"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": "Missing or malformed JWT", "data": nil})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"success": false, "message": "Invalid or expired JWT", "data": nil})
}

func jwtSuccess(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userId := token.Claims.(jwt.MapClaims)["id"]

	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)

	user, err := userService.FindOne(fmt.Sprintf("%s", userId))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	c.Locals("authUser", user)

	return c.Next()
}

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}
