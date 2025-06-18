package Routes

import (
	Handler "auth-api/Handler"
	"github.com/gofiber/fiber/v2"
)

func Database(app fiber.Router) {
	app.Post("/", Handler.Home)
}