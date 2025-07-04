package Middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Notfound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"code":    404,
		"message": "Not Found",
	})
}