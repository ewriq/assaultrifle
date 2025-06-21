package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"assaultrifle/Form"

	"github.com/gofiber/fiber/v2"
)


func ContainerAddHandler(c *fiber.Ctx) error {
	var req Form.ContainerAddRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid input",
		})
	}

	success, token := Database.ContainerAdd(req.Name, req.Password, req.Port, req.User, req.Type)
	if success {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Container oluşturuldu",
			"token":   token,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"status": "error",
		"error":  "Container oluşturulamadı",
	})
}


func ContainerListMyHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"`
	}
	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı adı eksik",
		})
	}

	container, err := Database.ContainerListMy(body.User)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   container,
	})
}

func ContainerListAllHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"`
	}
	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı adı eksik",
		})
	}

	containers, err := Database.ContainerListAll(body.User)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   containers,
	})
}

func ContainerDeleteHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	success, _ := Database.ContainerDelete(body.Token)
	if success {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Container silindi",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"status": "error",
		"error":  "Container silinemedi",
	})
}

func ContainerStartHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	if err := Container.StartContainer(body.Token); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  "Container başlatılamadı",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container başlatıldı",
	})
}

func ContainerStopHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	if err := Container.StopContainer(body.Token); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  "Container durdurulamadı",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container durdu",
	})
}
