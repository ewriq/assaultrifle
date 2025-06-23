package Routes

import (
	Handler "assaultrifle/Handler"

	"github.com/gofiber/fiber/v3"
)

func Auth(app fiber.Router) {
	app.Post("/login", Handler.Login)
	app.Post("/register", Handler.Register)
	app.Post("/user", Handler.User)
}
