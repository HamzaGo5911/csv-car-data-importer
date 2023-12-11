package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/HamzaGo5911/csv-car-data-importer/config"
	"github.com/HamzaGo5911/csv-car-data-importer/handlers"
	"github.com/HamzaGo5911/csv-car-data-importer/service"
)

func main() {
	err := service.GenerateCSV()
	if err != nil {
		log.Fatalf("Error generating CSV: %v", err)
	}

	config.ConnectToDb()
	go backgroundTask()

	r := gin.Default()

	carGroup := r.Group("/api")
	carGroup.GET("/cars", handlers.GetCars)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server shutting down...")
		log.Println("Background tasks finished. Exiting...")
		os.Exit(0)
	}()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func backgroundTask() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	importData := func() {
		err := service.ImportCSVToDatabase("cars.csv")
		if err != nil {
			log.Printf("Error importing CSV to database: %v", err)
		} else {
			log.Println("CSV imported successfully")
		}
	}

	importData()

	for {
		select {
		case <-ticker.C:
			importData()
		}
	}
}
