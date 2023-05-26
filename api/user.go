package api

import (
	"github.com/chensheep/hotel-reservation-backend/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Wei",
		LastName:  "Chen",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World " + c.Params("id"))
}
