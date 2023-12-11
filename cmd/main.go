package main

import (
	"github.com/HamzaGo5911/csv-car-data-importer/service"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/HamzaGo5911/csv-car-data-importer/config"
	"github.com/HamzaGo5911/csv-car-data-importer/db"
	"github.com/gin-gonic/gin"
)

var (
	wg sync.WaitGroup
)

func main() {
	config.ConnectToDb()
	go backgroundTask()

	r := gin.Default()

	carGroup := r.Group("/api")
	carGroup.GET("/cars", db.GetCars)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server shutting down...")

		wg.Wait() // Wait for background goroutines to finish
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

	wg.Add(1)
	defer wg.Done()

	importData := func() {
		err := service.ImportCSVToDatabase("cars.csv")
		if err != nil {
			log.Printf("Error importing CSV to database: %v", err)
		} else {
			log.Println("CSV imported successfully")
		}
	}

	importData() // Perform initial import on start

	for {
		select {
		case <-ticker.C:
			importData()
		}
	}
}
