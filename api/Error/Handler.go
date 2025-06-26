package Error

import (
	"assaultrifle/Form"

	"github.com/gofiber/fiber/v2"
)

func StatusBadRequest(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status": "error",
		"error":  "Geçersiz giriş",
	})
}

func InvalidPerm(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "error",
		"message": "Kullanıcı admin yetkisine sahip değil",
	})
}


func CustomError(c *fiber.Ctx, text string) error {
	return c.JSON(fiber.Map{
		"status":  "error",
		"message": text,
	})
}

func CustomSuccess(c *fiber.Ctx, text string) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": text,
	})
}
func CustomSuccessContainer(c *fiber.Ctx, data[]Form.Container, text string) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": text,
		"data": data,
	})
}

func CustomSuccessAuth(c *fiber.Ctx, data, text string) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": text,
		"data": data,
	})
}

func CustomSuccessUser(c *fiber.Ctx, data[]Form.User, text string) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": text,
		"data": data,
	})
}

func CustomSuccessContainerLog(c *fiber.Ctx, data, text string) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": text,
		"data": data,
	})
}



func CustomSuccessStatus(c *fiber.Ctx,  data interface{}, text string) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": text,
		"data": data,
	})
}
