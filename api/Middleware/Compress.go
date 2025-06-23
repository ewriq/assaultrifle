package Middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Compress(c *fiber.Ctx) error {
	compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
	return c.Next()
}
