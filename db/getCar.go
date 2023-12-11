package db

import (
	"github.com/HamzaGo5911/csv-car-data-importer/config"
	"github.com/HamzaGo5911/csv-car-data-importer/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetCars(c *gin.Context) {
	var cars []models.Car

	if err := config.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching cars",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}
