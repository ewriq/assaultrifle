package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"assaultrifle/Form"
	"assaultrifle/Utils"

	"github.com/gofiber/fiber/v2"
)


func ContainerAddHandler(c *fiber.Ctx) error {
	var req Form.ContainerAddRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	success, token := Database.ContainerAdd(Utils.Encode(req.Name) , Utils.Encode(req.Password), Utils.Encode(req.Port), Utils.Encode(req.User), Utils.Encode(req.Type))
	if success {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Container oluşturuldu",
			"token":   token,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status": "error",
		"error":  "Container oluşturulamadı",
	})
}


func ContainerListMyHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"`
	}

	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı adı eksik",
		})
	}


	container, err := Database.ContainerListMy(body.User)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı adı eksik",
		})
	}


	containers, err := Database.ContainerListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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


	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status": "error",
		"error":  "Container silinemedi",
	})
}


func ContainerStartHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	if err := Container.StartContainer(body.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Container başlatılamadı",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container başlatıldı",
	})
}

func ContainerRestartHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	if err := Container.RestartContainer(body.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Container başlatılamadı",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container başlatıldı",
	})
}


func ContainerLogsHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	 data, err := Container.GetContainerLogs(body.Token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Container logları alınamadı",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container logları alındı",
		"data": Utils.Encode(data),
	})
}

func ContainerStopHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	if err := Container.StopContainer(body.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Container durdurulamadı",
		})
	}


	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container durdu",
	})
}


func ContainerStatusHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	data, err := Container.GetContainerStatus(body.Token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Veri alınamadı",
		})
	}


	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   data,
	})
}
