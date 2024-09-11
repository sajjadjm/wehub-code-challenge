package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sajjadjm/wehub-code-challenge/internal/adapters/db"
	"github.com/sajjadjm/wehub-code-challenge/internal/adapters/io"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/services"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var csvFilePath string
var workerCount int

// main function sets up the root command for the CLI
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var rootCmd = &cobra.Command{
		Use:   "csv_processor",
		Short: "A CLI tool to process CSV and store data in MongoDB",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Starting CSV processor with the following configuration:\n")
			fmt.Printf("CSV File: %s\n", csvFilePath)
			fmt.Printf("Worker Count: %d\n", workerCount)

			// Setup MongoDB and CSV reader
			mongoRepo, err := db.NewCSVRecordRepository(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DATABASE"), "csv_data")
			if err != nil {
				log.Fatal("Failed to connect to MongoDB:", err)
			}

			csvReader := io.NewCSVReader()

			// Setup CSV processor with the provided worker count
			csvProcessor := services.NewCSVRecordService(csvReader, mongoRepo, workerCount)

			// Start processing the CSV file
			err = csvProcessor.ProcessCSV(csvFilePath)
			if err != nil {
				log.Fatal("Failed to process CSV:", err)
			}

			log.Println("CSV processing completed!")
		},
	}

	// Add flags for the CLI
	rootCmd.Flags().StringVarP(&csvFilePath, "csv", "c", "", "Path to the CSV file (required)")
	rootCmd.Flags().IntVarP(&workerCount, "workers", "w", 5, "Number of concurrent workers")

	// Mark the 'csv' flag as required
	if err := rootCmd.MarkFlagRequired("csv"); err != nil {
		fmt.Println(err)
	}

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
