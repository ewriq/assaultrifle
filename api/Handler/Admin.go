package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"assaultrifle/Form"

	"github.com/gofiber/fiber/v2"
)

func Shutdown(c *fiber.Ctx) error {
	var reqbody Form.UserInfo

	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	perm, _ := Database.ValidateAdminAccess(reqbody.Token)
	if perm == "user" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Kullanıcı admin yetkisine sahip değil",
		})
	}


	stop := Container.ContainerStopAll()
	if stop != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Durdurulamadı",
		})
	}
	
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Konteynerlerin tümü durduruldu",
	})
}