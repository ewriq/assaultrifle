package Routes

import (
	Handler "assaultrifle/Handler"

	"github.com/gofiber/fiber/v2"
)

func Admin(app fiber.Router) {
	app.Post("/shutdown", Handler.Shutdown)
	app.Post("/del", Handler.DeleteAllContainer)
	app.Post("/list", Handler.ListAllContainer)
	app.Get("/log", Handler.GetLog)
}