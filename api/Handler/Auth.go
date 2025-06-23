package Handler

import (
	"assaultrifle/Database"
	"assaultrifle/Form"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Login(c fiber.Ctx) error {
	var reqbody Form.UserBody

	if err := c.Bind().Body(&reqbody); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid input",
		})
	}

	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid email or password format",
		})
	}

	token, err := Database.Login(reqbody.Email, reqbody.Password)
	if err != nil {
		return c.Status(502).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "User logged in successfully",
		"token":   token,
	})
}

func Register(c fiber.Ctx) error {
	var reqbody Form.UserBody

	if err := c.Bind().Body(&reqbody); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid input",
		})
	}

	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid email or password format",
		})
	}

	success, token := Database.Register(reqbody.Email, reqbody.Password, reqbody.Username)
	if !success {
		return c.Status(502).JSON(fiber.Map{
			"status": "error",
			"error":  "Registration failed",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "User registered successfully",
		"token":   token,
	})
}

func User(c fiber.Ctx) error {
	var reqbody Form.UserInfo
	if err := c.Bind().Body(&reqbody); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid input",
		})
	}

	user, err := Database.Users(reqbody.Token)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "User not found",
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
