package api

import (
	"github.com/Andre6-dev/hotel-reservation-api/models"
	"github.com/gofiber/fiber/v2"
)

// UserHandler is a handler for the /api/v1/user route
func HandlerListUser(c *fiber.Ctx) error {
	u := models.User{
		FirstName: "Andre",
		LastName:  "Silva",
	}
	return c.JSON(u)
}

// UserHandler is a handler for the /api/v1/user/:id route
func HandlerGetUser(c *fiber.Ctx) error {
	return c.JSON("Andre")
}
