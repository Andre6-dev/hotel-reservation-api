package api

import (
	"context"

	"github.com/Andre6-dev/hotel-reservation-api/db"
	"github.com/Andre6-dev/hotel-reservation-api/models"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

// Constructor for UserHandler
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// UserHandler is a handler for the /api/v1/user/:id route
func (h *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

// UserHandler is a handler for the /api/v1/user route
func (h *UserHandler) HandlerListUsers(c *fiber.Ctx) error {
	u := models.User{
		FirstName: "Andre",
		LastName:  "Silva",
	}
	return c.JSON(u)
}
