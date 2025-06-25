package Routes

import (
	Handler "assaultrifle/Handler"

	"github.com/gofiber/fiber/v2"
)

func Admin(app fiber.Router) {
	app.Post("/shutdown", Handler.Shutdown)
	app.Post("/register", Handler.Register)
	app.Post("/user", Handler.User)
}