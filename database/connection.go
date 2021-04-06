package database

import (
	"Project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root1234@/polls"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Options{})
	connection.AutoMigrate(&models.Polls{})
	connection.AutoMigrate(&models.OptionUser{})
	connection.AutoMigrate(&models.UserVoteQues{})
}
