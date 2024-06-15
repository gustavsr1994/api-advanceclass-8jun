package config

import (
	"example/api-advance-class/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@dm1n1994@tcp(localhost:3306)/cource_net"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.Product{})
	DB = database
}
