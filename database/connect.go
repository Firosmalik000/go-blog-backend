package database

import (
	"backend-project/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	database, err :=  gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println(("Connected to MySql database"))
	}
	DB=database
	// AutoMigrate untuk connect ke database dan membuat table baru berdasarkan struct models
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}