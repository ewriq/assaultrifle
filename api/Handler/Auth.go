package Handler

import (
	"assaultrifle/Database"
	"assaultrifle/Form"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var reqbody Form.UserBody
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	if isValidEmail(reqbody.Email) && isPasswordValid(reqbody.Password) {
		token, err := Database.Login(reqbody.Email, reqbody.Password)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Giriş başarılı",
			"token":   token,
		})
	}

	return c.Status(400).JSON(fiber.Map{
		"status": "error",
		"error":  "Geçersiz email ya da şifre",
	})
}

func Register(c *fiber.Ctx) error {
	var reqbody Form.UserBody
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	if isValidEmail(reqbody.Email) && isPasswordValid(reqbody.Password) {
		ok, token := Database.Register(reqbody.Email, reqbody.Password, reqbody.Username)
		if !ok {
			return c.Status(400).JSON(fiber.Map{
				"status": "error",
				"error":  "Email zaten kayıtlı veya kayıt hatası",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Kayıt başarılı",
			"token":   token,
		})
	}

	return c.Status(400).JSON(fiber.Map{
		"status": "error",
		"error":  "Geçersiz parola veya mail",
	})
}

func User(c *fiber.Ctx) error {
	var reqbody Form.UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	user, err := Database.Users(reqbody.Token)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})
}


func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func isPasswordValid(password string) bool {
	return len(password) >= 8
}

