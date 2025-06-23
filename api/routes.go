package main

import (
	Handler "assaultrifle/Handler"
	Middleware "assaultrifle/Middleware"
	"assaultrifle/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Initalize(app *fiber.App) {
	app.Use(Middleware.Cors)
	app.Use(Middleware.RateLimit)
	app.Use(helmet.New())
	app.Use(Middleware.Compress)

	auth := app.Group("/api/auth/")
	container := app.Group("/api/container/")
	
	app.Get("/", Handler.Home)
	app.Get("/ws/:id", websocket.New(Handler.WebSocket))

	Routes.Auth(auth)
	Routes.Container(container)

	app.Use(Middleware.Notfound)
}
