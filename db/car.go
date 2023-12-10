package db

import (
	"csv-car-data-importer/config"
	"csv-car-data-importer/models"
	"fmt"
)

// SaveData saves the car data in the database
func SaveData(car models.Car) error {
	err := config.DB.Create(&car).Error
	if err != nil {
		return fmt.Errorf("failed to save car data: %v", err)
	}
	return nil
}
