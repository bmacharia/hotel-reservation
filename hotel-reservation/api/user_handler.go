package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hotel-reservation/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Babu",
		LastName:  "Doube Fisting",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Babu")
}
