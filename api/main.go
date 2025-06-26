package main

import (
	"assaultrifle/Handler"
	"assaultrifle/Middleware"
	"assaultrifle/Routes"


	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})


	Initialize(app)

	app.Listen(":3000")
}


func Initialize(app *fiber.App) {

	app.Use(Middleware.Cors)
	app.Use(Middleware.RateLimit)
	app.Use(helmet.New())
	app.Use(Middleware.Compress)

	auth := app.Group("/api/auth")
	container := app.Group("/api/container")
	admin := app.Group("/api/admin")


	app.Get("/", Handler.Home)


	app.Get("/ws/:id", websocket.New(Handler.WebSocket))


	Routes.Auth(auth)
	Routes.Admin(admin)
	Routes.Container(container)

	app.Use(Middleware.Notfound)
}
