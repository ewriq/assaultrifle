package Form

import "gorm.io/gorm"

type ContainerAddRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Type     string `json:"type"`
}

type Container struct {
	gorm.Model
	Name     string
	Password string
	Port     string `gorm:"unique"`
	Token    string `gorm:"unique"`
	User     string
	Type     string
	ID       string
}

func (Container) TableName() string {
	return "container"
}