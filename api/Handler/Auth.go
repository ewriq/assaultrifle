package Handler

import (
	"assaultrifle/Database"
	"assaultrifle/Form"
	"assaultrifle/Utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)


func Login(c *fiber.Ctx) error {
	var reqbody Form.UserBody


	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}


	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz e-posta veya şifre formatı",
		})
	}


	token, err := Database.Login(Utils.Encode(reqbody.Email) , Utils.Encode(reqbody.Password))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}


	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Kullanıcı başarıyla giriş yaptı",
		"token":   token,
	})
}


func Register(c *fiber.Ctx) error {
	var reqbody Form.UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz e-posta veya şifre formatı",
		})
	}


	success, token := Database.Register(Utils.Encode(reqbody.Email), Utils.Encode(reqbody.Password), Utils.Encode(reqbody.Username))
	if !success {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "error",
			"error":  "Kayıt başarısız",
		})
	}


	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Kullanıcı başarıyla kaydedildi",
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	var reqbody Form.UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}


	user, err := Database.Users(reqbody.Token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı bulunamadı",
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
