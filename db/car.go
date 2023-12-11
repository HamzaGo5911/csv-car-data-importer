package db

import (
	"fmt"

	"github.com/HamzaGo5911/csv-car-data-importer/config"
	"github.com/HamzaGo5911/csv-car-data-importer/models"
)

// SaveData saves the car data in the database
func SaveData(car models.Car) error {
	err := config.DB.Create(&car).Error
	if err != nil {
		return fmt.Errorf("failed to save car data: %v", err)
	}
	return nil
}
