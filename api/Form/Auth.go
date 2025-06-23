package Form

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string `gorm:"unique"`
	Token    string `gorm:"unique"`
	Perm     string
}


type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Token string `json:"token"`
}


func (User) TableName() string {
	return "user"
}
