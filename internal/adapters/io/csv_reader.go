package io

import (
	"encoding/csv"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/domain"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/ports"
	"log"
	"os"
	"strconv"
)

type CSVReader struct{}

// NewCSVReader creates a new CSVReader
func NewCSVReader() ports.CSVReaderInterface {
	return &CSVReader{}
}

// ReadCSV reads the CSV file and returns the data as a list of CSVRecords.
func (r *CSVReader) ReadCSV(filePath string) ([]domain.CSVRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error in closing file: ", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var csvRecords []domain.CSVRecord
	for _, record := range records {
		magnitude, _ := strconv.Atoi(record[6])
		dataValue, _ := strconv.ParseFloat(record[2], 64)

		csvRecord := domain.CSVRecord{
			SeriesReference: record[0],
			Period:          record[1],
			DataValue:       dataValue,
			Suppressed:      record[3],
			Status:          record[4],
			Units:           record[5],
			Magnitude:       magnitude,
			Subject:         record[7],
			Group:           record[8],
			SeriesTitle1:    record[9],
			SeriesTitle2:    record[10],
			SeriesTitle3:    record[11],
			SeriesTitle4:    record[12],
			SeriesTitle5:    record[13],
		}
		csvRecords = append(csvRecords, csvRecord)
	}
	return csvRecords, nil
}
