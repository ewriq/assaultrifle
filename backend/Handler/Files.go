package Handler

import (
	"assaultrifle/Container"
	"assaultrifle/Database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ContainerListFilesHandler(c *fiber.Ctx) error {
	var body struct {
		Token     string `json:"token"`
		Path      string `json:"path"`
		User string `json:"user"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" || body.Path == "" || body.User == "" {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token, path veya kullanıcı eksik",
		})
	}
	ok, err := Database.ValidateContainerAccess(body.Token, body.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	files, err := Container.ListContainerFiles(body.Token, body.Path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   files,
	})
}


func ContainerDeleteFileHandler(c *fiber.Ctx) error {
	var body struct {
		Token     string `json:"token"`
		Path      string `json:"path"`
		UserToken string `json:"user"`
	}

	if err := c.BodyParser(&body); err != nil || body.Token == "" || body.Path == "" || body.UserToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token, dosya yolu veya kullanıcı eksik",
		})
	}

	ok, err := Database.ValidateContainerAccess(body.Token, body.UserToken)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	fmt.Println(body.Path)
	if err := Container.DeleteContainerFile(body.Token, body.Path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Dosya silindi",
	})
}


func ContainerUploadFileHandler(c *fiber.Ctx) error {
	containerID := c.FormValue("token")
	targetPath := c.FormValue("target")
	userToken := c.FormValue("user")

	if containerID == "" || targetPath == "" || userToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Token, hedef yol veya kullanıcı eksik",
		})
	}

	ok, err := Database.ValidateContainerAccess(containerID, userToken)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Dosya alınamadı",
		})
	}

	savePath := fmt.Sprintf("./Tmp/%s", file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Dosya kaydedilemedi",
		})
	}

	if err := Container.CopyFileToContainer(containerID, savePath, targetPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Dosya yüklendi",
	})
}
