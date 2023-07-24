package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/playground/h-reservation/db"
	"github.com/playground/h-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users := types.User{
		// ID:        "123",
		FirstName: "Ko Sai",
		LastName:  "Hlaing Lu",
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id      = c.Params("id")
		context = context.Background()
	)

	user, err := h.userStore.GetUserById(context, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
