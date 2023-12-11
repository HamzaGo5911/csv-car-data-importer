package db

import (
	"github.com/HamzaGo5911/csv-car-data-importer/config"
	"github.com/HamzaGo5911/csv-car-data-importer/models"
)

// GetAllCars get all the car data from the database
func GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	if err := config.DB.Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}
