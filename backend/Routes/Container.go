package Routes

import (
	Handler "assaultrifle/Handler"

	"github.com/gofiber/fiber/v2"
)

func Container(app fiber.Router) {
	app.Post("/del", Handler.ContainerDeleteHandler)
	app.Post("/add", Handler.ContainerAddHandler)
	app.Post("/list", Handler.ContainerListMyHandler)
	app.Post("/listall", Handler.ContainerListAllHandler)
	app.Post("/start", Handler.ContainerStartHandler)
	app.Post("/stop", Handler.ContainerStopHandler)
	app.Post("/status", Handler.ContainerStatusHandler)
	app.Post("/restart", Handler.ContainerRestartHandler)
	app.Post("/log", Handler.ContainerLogsHandler)
	app.Post("/file/add", Handler.ContainerUploadFileHandler)
	app.Post("/file/del", Handler.ContainerDeleteFileHandler)
	app.Post("/file/list", Handler.ContainerListFilesHandler)	
}

/*
del
add
listmy
listall
start
stop
*/
