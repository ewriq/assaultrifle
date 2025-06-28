package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"assaultrifle/Error"
	"assaultrifle/Form"

	"github.com/gofiber/fiber/v2"
)

func Shutdown(c *fiber.Ctx) error {
	var reqbody Form.UserInfo

	if err := c.BodyParser(&reqbody); err != nil {
		return Error.StatusBadRequest(c)
	}

	perm, _ := Database.ValidateAdminAccess(reqbody.Token)
	if perm == "user" {
		return Error.InvalidPerm(c)
	}


	stop := Container.ContainerStopAll()
	if stop != nil {
		return 	Error.CustomError(c, "Konteynerler durdurulamadı.")
	}
	
	return Error.CustomSuccess(c, "Tüm konteynerler başarıyla durduruldu.")
}

func DeleteAllContainer(c *fiber.Ctx) error {
	var reqbody Form.UserInfo

	if err := c.BodyParser(&reqbody); err != nil {
		return Error.StatusBadRequest(c)
	}

	perm, _ := Database.ValidateAdminAccess(reqbody.Token)
	if perm == "user" {
		return Error.InvalidPerm(c)
	}


	del := Container.DeleteAllContainer()
	if del != nil {
		return 	Error.CustomError(c, "Konteynerler silinemedi.")
	}
	
	return Error.CustomSuccess(c, "Tüm konteynerler başarıyla silindi.")
}

func ListAllContainer(c *fiber.Ctx) error {
	var reqbody Form.UserInfo

	if err := c.BodyParser(&reqbody); err != nil {
		return Error.StatusBadRequest(c)
	}

	perm, _ := Database.ValidateAdminAccess(reqbody.Token)
	if perm == "user" {
		return Error.InvalidPerm(c)
	}


	data, err := Database.ContainerListAll()
	if err != nil {
		return 	Error.CustomError(c, "Konteynerler listelenemedi.")
	}
	
	return Error.CustomSuccessContainer(c,data,"Konteynerlerin tümü listelendi.")
}