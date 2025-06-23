package Database

import (
	"assaultrifle/Form"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root@tcp(localhost:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı: " + err.Error())
	}

	DB.AutoMigrate(&Form.User{}, &Form.Container{})
	fmt.Println("✅ GORM ile MySQL bağlantısı kuruldu ve tablolar migrat edildi.")
}
