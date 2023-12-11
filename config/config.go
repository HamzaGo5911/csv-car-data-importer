package config

import (
	"fmt"
	"github.com/HamzaGo5911/csv-car-data-importer/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := os.Getenv("DSN_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	err = DB.AutoMigrate(&models.Car{})
	if err != nil {
		panic("Cannot auto migrate tables")
	}
}
