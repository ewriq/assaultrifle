package Routes

import (
	Handler "assaultrifle/Handler"

	"github.com/gofiber/fiber/v2"
)

func Auth(app fiber.Router) {
	app.Post("/login", Handler.Login)
	app.Post("/register", Handler.Register)
	app.Post("/user", Handler.User)
	app.Post("/settings/reset/password", Handler.User)
	app.Post("/settings/reset/email", Handler.User)
	app.Post("/settings/reset/username", Handler.User)
	app.Post("/verify/email", Handler.User)
	app.Post("/device/list", Handler.User)
}