package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/HamzaGo5911/csv-car-data-importer/db"
	"github.com/HamzaGo5911/csv-car-data-importer/models"
)

func ImportCSVToDatabase(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV: %v", err)
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		year, _ := strconv.Atoi(record[2])
		sellingPrice, _ := strconv.ParseFloat(record[3], 64)

		car := models.Car{
			Name:         record[1],
			Year:         year,
			SellingPrice: sellingPrice,
			Transmission: record[4],
		}

		if err := db.SaveData(car); err != nil {
			fmt.Printf("Error saving car data: %v\n", err)
		} else {
			fmt.Println("Car data saved successfully")
		}
	}

	return nil
}
