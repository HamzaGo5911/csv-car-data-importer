package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/HamzaGo5911/csv-car-data-importer/models"
)

// GenerateCSV generates a CSV file with sample car data
func GenerateCSV() error {
	file, err := os.Create("cars.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"id", "name", "year", "selling_price", "transmission"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for i := 1; i < 50; i++ {
		car := generateSampleCarData(i)
		row := []string{
			strconv.Itoa(car.ID),
			car.Name,
			strconv.Itoa(car.Year),
			strconv.FormatFloat(car.SellingPrice, 'f', 2, 64),
			car.Transmission,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	log.Println("CSV file generated successfully")
	return nil
}

func generateSampleCarData(i int) models.Car {
	return models.Car{
		ID:           i,
		Name:         "Car_" + strconv.Itoa(i),
		Year:         2000 + i%20,
		SellingPrice: float64(10000 + i*1000),
		Transmission: "Automatic",
	}
}
