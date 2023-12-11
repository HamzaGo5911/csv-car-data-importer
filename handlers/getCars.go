package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/HamzaGo5911/csv-car-data-importer/db"
	"github.com/HamzaGo5911/csv-car-data-importer/models"
)

// GetCars handle cars request
func GetCars(c *gin.Context) {
	var cars []models.Car

	cars, err := db.GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching cars",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}
