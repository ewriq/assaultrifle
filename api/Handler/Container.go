package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"assaultrifle/Form"

	// Fiber v2 import yolu
	"github.com/gofiber/fiber/v2"
)

// ContainerAddHandler yeni bir container ekler
func ContainerAddHandler(c *fiber.Ctx) error {
	var req Form.ContainerAddRequest
	// İstek gövdesini ContainerAddRequest yapısına bağlar (Fiber v2 için BodyParser)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	// Veritabanı üzerinden container ekler
	success, token := Database.ContainerAdd(req.Name, req.Password, req.Port, req.User, req.Type)
	if success {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Container oluşturuldu",
			"token":   token,
		})
	}

	// Container ekleme başarısız olursa hata döner
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status": "error",
		"error":  "Container oluşturulamadı",
	})
}

// ContainerListMyHandler belirli bir kullanıcıya ait container'ları listeler
func ContainerListMyHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"`
	}
	// İstek gövdesini bağlar ve kullanıcı adının boş olmadığını kontrol eder (Fiber v2 için BodyParser)
	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı adı eksik",
		})
	}

	// Veritabanından kullanıcının container'larını listeler
	container, err := Database.ContainerListMy(body.User)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Container listesiyle başarılı yanıt döner
	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   container,
	})
}

// ContainerListAllHandler tüm container'ları listeler (yönetici yetkisi gerektirebilir)
func ContainerListAllHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"` // Kullanıcı bilgisi burada neden isteniyor? Yetkilendirme için mi?
	}
	// İstek gövdesini bağlar ve kullanıcı adının boş olmadığını kontrol eder (Fiber v2 için BodyParser)
	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı adı eksik",
		})
	}

	// Veritabanından tüm container'ları listeler
	containers, err := Database.ContainerListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Tüm container'larla başarılı yanıt döner
	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   containers,
	})
}

// ContainerDeleteHandler bir container'ı siler
func ContainerDeleteHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	// İstek gövdesini bağlar ve token'ın boş olmadığını kontrol eder (Fiber v2 için BodyParser)
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	// Veritabanı üzerinden container'ı siler
	success, _ := Database.ContainerDelete(body.Token)
	if success {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Container silindi",
		})
	}

	// Container silme başarısız olursa hata döner
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status": "error",
		"error":  "Container silinemedi",
	})
}

// ContainerStartHandler bir container'ı başlatır
func ContainerStartHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	// İstek gövdesini bağlar ve token'ın boş olmadığını kontrol eder (Fiber v2 için BodyParser)
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	// Container'ı başlatır
	if err := Container.StartContainer(body.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Container başlatılamadı",
		})
	}

	// Başarılı başlatma yanıtı döner
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container başlatıldı",
	})
}

// ContainerStopHandler bir container'ı durdurur
func ContainerStopHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	// İstek gövdesini bağlar ve token'ın boş olmadığını kontrol eder (Fiber v2 için BodyParser)
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	// Container'ı durdurur
	if err := Container.StopContainer(body.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Container durdurulamadı",
		})
	}

	// Başarılı durdurma yanıtı döner
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Container durdu",
	})
}

// ContainerStatusHandler bir container'ın durumunu getirir
func ContainerStatusHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	// İstek gövdesini bağlar ve token'ın boş olmadığını kontrol eder (Fiber v2 için BodyParser)
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token eksik",
		})
	}

	// Container'ın durumunu alır
	data, err := Container.GetContainerStatus(body.Token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Veri alınamadı",
		})
	}

	// Container durumuyla başarılı yanıt döner
	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   data,
	})
}
