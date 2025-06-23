
package Handler

import (
	"assaultrifle/Database"
	"assaultrifle/Form"
	"strings"

	// Fiber v2 import yolu
	"github.com/gofiber/fiber/v2"
	// WebSocket handler'ı Home ile birlikte olmadığı için burada websocket'i kaldırdık.
	// Eğer Home veya Login/Register/User içinde websocket kullanılıyorsa tekrar eklenmeli.
)

// Home handler'ı ana sayfa isteğini karşılar


// Login handler'ı kullanıcı girişi yapar
func Login(c *fiber.Ctx) error {
	var reqbody Form.UserBody

	// İstek gövdesini UserBody yapısına bağlar (Fiber v2 için BodyParser)
	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	// E-posta ve şifre formatlarını doğrular
	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz e-posta veya şifre formatı",
		})
	}

	// Veritabanı üzerinden giriş yapar ve token alır
	token, err := Database.Login(reqbody.Email, reqbody.Password)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Başarılı giriş yanıtı döner
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Kullanıcı başarıyla giriş yaptı",
		"token":   token,
	})
}

// Register handler'ı yeni kullanıcı kaydı yapar
func Register(c *fiber.Ctx) error {
	var reqbody Form.UserBody

	// İstek gövdesini UserBody yapısına bağlar (Fiber v2 için BodyParser)
	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	// E-posta ve şifre formatlarını doğrular
	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz e-posta veya şifre formatı",
		})
	}

	// Veritabanı üzerinden kayıt yapar ve token alır
	success, token := Database.Register(reqbody.Email, reqbody.Password, reqbody.Username)
	if !success {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "error",
			"error":  "Kayıt başarısız",
		})
	}

	// Başarılı kayıt yanıtı döner
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Kullanıcı başarıyla kaydedildi",
		"token":   token,
	})
}

// User handler'ı kullanıcı bilgilerini döner
func User(c *fiber.Ctx) error {
	var reqbody Form.UserInfo
	// İstek gövdesini UserInfo yapısına bağlar (Fiber v2 için BodyParser)
	if err := c.BodyParser(&reqbody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Geçersiz giriş",
		})
	}

	// Veritabanından kullanıcı bilgilerini token ile alır
	user, err := Database.Users(reqbody.Token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Kullanıcı bulunamadı",
		})
	}

	// Kullanıcı bilgileriyle başarılı yanıt döner
	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})
}

// isValidEmail basit bir e-posta formatı doğrulaması yapar (içerisinde '@' kontrolü)
func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

// isPasswordValid basit bir şifre uzunluğu doğrulaması yapar (en az 8 karakter)
func isPasswordValid(password string) bool {
	return len(password) >= 8
}

// WebSocket handler'ı, WebSocket bağlantılarını yönetir
// NOT: Bu handler genellikle ayrı bir dosyada veya webSocketRoute'ta tanımlanır.
// Ancak Home ile aynı dosyada olmasından dolayı burada bırakıldı.
// Eğer sadece webSocket için kullanılıyorsa, "github.com/gofiber/websocket/v2" import'u gereklidir.
// Örneğin: "assaultrifle/main.go" dosyasındaki app.Get("/ws/:id", websocket.New(Handler.WebSocket)) çağrısı için gereklidir.
// Bu dosya içerisinde kullanılmadığı için import'u kaldırılmıştır.
// Eğer kullanılacaksa, aşağıdaki yorum satırını kaldırın:
// import "github.com/gofiber/websocket/v2"
// Bu fonksiyonu ana dosyanıza geri taşıdık
/*
func WebSocket(c *fiber.Ctx) error {
	// WebSocket bağlantısını yükseltir ve id parametresini kullanır
	return c.SendString("WebSocket endpoint hit for ID: " + c.Params("id"))
}
*/