package main

import (
	"csv-car-data-importer/config"
	"csv-car-data-importer/db"
	"csv-car-data-importer/models"
	"csv-car-data-importer/service"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	config.ConnectToDb()

	err := service.GenerateCSV()
	if err != nil {
		log.Fatalf("Error generating CSV: %s", err)
	}

	filePath := "cars.csv"
	err = service.ImportCSVToDatabase(filePath)
	if err != nil {
		fmt.Println("Error importing CSV:", err)
		return
	}
	fmt.Println("CSV data imported successfully to the database")

	go func() {
		for {
			backgroundTask()
			time.Sleep(24 * time.Hour)
		}
	}()

	// Start HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the CSV Car Data Importer!")
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func backgroundTask() {
	var car = models.Car{} // This line might need to be adjusted to create meaningful sample data
	if err := db.SaveData(car); err != nil {
		log.Printf("Error saving car data: %s", err)
	} else {
		fmt.Println("Car data saved successfully")
	}
}
