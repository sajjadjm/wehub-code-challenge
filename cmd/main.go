package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sajjadjm/wehub-code-challenge/internal/adapters/db"
	"github.com/sajjadjm/wehub-code-challenge/internal/adapters/http"
	"github.com/sajjadjm/wehub-code-challenge/internal/adapters/io"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/services"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	repo, _ := db.NewCSVRecordRepository(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DATABASE"), "csv_data")
	csvReader := io.NewCSVReader()

	// Application service
	service := services.NewCSVRecordService(csvReader, repo, 10)

	// Fiber setup
	app := fiber.New()

	// Record handler
	recordHandler := http.NewCSVRecordHandler(service)

	// Routes
	app.Post("/records", recordHandler.CreateCSVRecord)
	app.Get("/records", recordHandler.GetAllCSVRecords)
	app.Get("/records/:id", recordHandler.GetCSVRecordByID)
	app.Put("/records/:id", recordHandler.UpdateCSVRecord)
	app.Delete("/records/:id", recordHandler.DeleteCSVRecord)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
