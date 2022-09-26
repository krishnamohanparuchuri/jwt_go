package database

import (
	"github.com/krishna/jwt_go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Krishna2022!!@/youtubegodb"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
