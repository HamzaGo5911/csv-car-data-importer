package db

import (
	"csv-car-data-importer/config"
	"csv-car-data-importer/models"
	"csv-car-data-importer/utils"
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

	utils.Paginate(c, len(cars), cars)
}
