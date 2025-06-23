package Database

import (
	"assaultrifle/Container"
	"assaultrifle/Form"
	"errors"
	"fmt"
)

func ContainerAdd(name, password, port, user, types string) (bool, string) {
	if FinderPort(port) {
		return false, ""
	}

	var image, token string
	var err error

	switch types {
	case "MySql":
		image = "mysql:8.0"
		Container.PullImage(image)
		token, err = Container.CreateMySQLContainer(name, image, port, password)
	case "PostgreSql":
		image = "postgres:16"
		Container.PullImage(image)
		token, err = Container.CreatePostgresContainer(name, image, port, password)
	case "NodeJS":
		image = "node:18"
		Container.PullImage(image)
		token, err = Container.CreateNodeContainer(name, image, port)
	case "Ubunutu":
		image = "ubuntu:22.04"
		Container.PullImage(image)
		token, err = Container.CreateNodeContainer(name, image, port)
	default:
		return false, ""
	}

	if err != nil {
		return false, ""
	}

	newContainer := Form.Container{
		Name:     name,
		Password: password,
		Port:     port,
		User:     user,
		Type:     types,
		Token:    token,
	}

	if err := DB.Create(&newContainer).Error; err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, token
}

func FinderPort(port string) bool {
	var container Form.Container
	err := DB.Where("port = ?", port).First(&container).Error
	return err == nil
}

func ContainerDelete(token string) (bool, string) {
	if err := Container.DeleteContainer(token); err != nil {
		fmt.Println("❌ Docker silme hatası:", err)
		return false, ""
	}

	if err := DB.Where("token = ?", token).Delete(&Form.Container{}).Error; err != nil {
		fmt.Println("❌ DB silme hatası:", err)
		return false, ""
	}

	return true, token
}

func ContainerListMy(user string) ([]Form.Container, error) {
	var containers []Form.Container
	if err := DB.Where("user = ?", user).Find(&containers).Error; err != nil {
		return nil, err
	}
	if len(containers) == 0 {
		return nil, errors.New("hiç container bulunamadı")
	}
	return containers, nil
}

func ContainerListAll() ([]Form.Container, error) {
	var containers []Form.Container
	if err := DB.Find(&containers).Error; err != nil {
		return nil, err
	}
	if len(containers) == 0 {
		return nil, errors.New("container bulunamadı")
	}
	return containers, nil
}
