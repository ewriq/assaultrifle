package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"assaultrifle/Error"
	"assaultrifle/Form"
	

	"github.com/gofiber/fiber/v2"
)


func ContainerAddHandler(c *fiber.Ctx) error {
	var req Form.ContainerAddRequest

	if err := c.BodyParser(&req); err != nil {
		return Error.StatusBadRequest(c)
	}

	success, _ := Database.ContainerAdd(req.Name , req.Password, req.Port, req.User,req.Type)
	if success {
		return Error.CustomSuccess(c,"Container oluşturuldu")
	}

	return Error.CustomError(c, "Container oluşturulamadı")
}


func ContainerListMyHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"`
	}

	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return Error.CustomError(c,  "Kullanıcı adı eksik")
	}

	container, err := Database.ContainerListMy(body.User)
	if err != nil {
		return Error.CustomError(c,err.Error())
	}

	return Error.CustomSuccessContainer(c, container, "Başarıyla konteynerler listelendi.")
}

func ContainerListAllHandler(c *fiber.Ctx) error {
	var body struct {
		User string `json:"user"` 
	}

	if err := c.BodyParser(&body); err != nil || body.User == "" {
		return Error.CustomError(c,  "Kullanıcı tokeni eksik")
	}

	containers, err := Database.ContainerListAll()
	if err != nil {
		return Error.CustomError(c, err.Error())
	}

	return Error.CustomSuccessContainer(c, containers, "Başarıyla konteynerler listelendi.")
}


func ContainerDeleteHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return Error.CustomError(c,"Token eksik")
	}

	success, _ := Database.ContainerDelete(body.Token)
	if success {
		return Error.CustomSuccess(c,"Container silindi" )
	}

	return Error.CustomError(c, "Container silinemedi")
}


func ContainerStartHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return Error.CustomError(c,"Token eksik")
	}

	if err := Container.StartContainer(body.Token); err != nil {
		return Error.CustomError(c, "Container başlatılamadı")
	}

	return Error.CustomSuccess(c,"Container başltatıldı" )
}

func ContainerRestartHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return Error.CustomError(c,"Token eksik")
	}

	if err := Container.RestartContainer(body.Token); err != nil {
		return Error.CustomError(c, "Container başlatılamadı")
	}

	return Error.CustomSuccess(c,"Container başltatıldı" )
}


func ContainerLogsHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return Error.CustomError(c,"Token eksik")
	}

	 data, err := Container.GetContainerLogs(body.Token)
	if err != nil {
		return Error.CustomError(c, "Container logları alınamadı")
	}

	return Error.CustomSuccessContainerLog(c,data,"Container logları alındı")
}

func ContainerStopHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	
	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		
	}

	if err := Container.StopContainer(body.Token); err != nil {
		return Error.CustomError(c, "Container durdurulamadı")
	}


	return Error.CustomSuccess(c,"Container durdu")
}


func ContainerStatusHandler(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" {
		return Error.CustomError(c,"Token eksik")
	}

	data, err := Container.GetContainerStatus(body.Token)
	if err != nil {
		return Error.CustomError(c,"Veri alınamadı")
	}

	return Error.CustomSuccessStatus(c, data, "Anlık durum listelendi.")
}
