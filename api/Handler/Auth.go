package Handler

import (
	"assaultrifle/Database"
	"assaultrifle/Error"
	"assaultrifle/Form"
	"strings"

	"github.com/gofiber/fiber/v2"
)


func Login(c *fiber.Ctx) error {
	var reqbody Form.UserBody


	if err := c.BodyParser(&reqbody); err != nil {
		return Error.StatusBadRequest(c)
	}


	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return Error.CustomError(c	,"Geçersiz e-posta veya şifre formatı")
	}


	token, err := Database.Login(reqbody.Email , reqbody.Password)
	if err != nil {
		return Error.CustomError(c	,err.Error())
	}


	return Error.CustomSuccessAuth(c,token,"Kullanıcı başarıyla giriş yaptı")
}


func Register(c *fiber.Ctx) error {
	var reqbody Form.UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return Error.StatusBadRequest(c)
	}

	if !isValidEmail(reqbody.Email) || !isPasswordValid(reqbody.Password) {
		return Error.CustomError(c	,"Geçersiz e-posta veya şifre formatı")
	}


	success, token := Database.Register(reqbody.Email, reqbody.Password, reqbody.Username)
	if !success {
		return Error.CustomError(c	, "Kayıt başarısız")
	}


	return Error.CustomSuccessAuth(c,token,"Kullanıcı başarıyla kaydedildi")
}

func User(c *fiber.Ctx) error {
	var reqbody Form.UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return Error.StatusBadRequest(c)
	}


	user, err := Database.Users(reqbody.Token)
	if err != nil {
		return Error.CustomError(c	,"Kullanıcı bulunamadı")
	}

	return Error.CustomSuccessUser(c, []Form.User{user},	"Kullanıcı bilgileri başarıyla çekildi.")
}


func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}


func isPasswordValid(password string) bool {
	return len(password) >= 8
}
