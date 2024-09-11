package services

import (
	"github.com/sajjadjm/wehub-code-challenge/internal/core/domain"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/ports"
	"log"
	"sync"
)

type CSVRecordService struct {
	csvReader   ports.CSVReaderInterface
	repo        ports.CSVRecordRepositoryInterface
	workerCount int
}

func NewCSVRecordService(csvReader ports.CSVReaderInterface, repo ports.CSVRecordRepositoryInterface, workerCount int) *CSVRecordService {
	return &CSVRecordService{
		csvReader:   csvReader,
		repo:        repo,
		workerCount: workerCount,
	}
}

func (s *CSVRecordService) ProcessCSV(filePath string) error {
	records, err := s.csvReader.ReadCSV(filePath)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	recordChannel := make(chan domain.CSVRecord)

	for i := 0; i < s.workerCount; i++ {
		wg.Add(1)
		go s.processRecords(recordChannel, &wg, i)
	}

	go func() {
		defer close(recordChannel)
		for _, record := range records {
			recordChannel <- record
		}
	}()

	wg.Wait()
	return nil
}

func (s *CSVRecordService) CreateCSVRecord(record *domain.CSVRecord) error {
	return s.repo.Create(record)
}

func (s *CSVRecordService) GetCSVRecordByID(id string) (*domain.CSVRecord, error) {
	return s.repo.GetByID(id)
}

func (s *CSVRecordService) UpdateCSVRecord(id string, record *domain.CSVRecord) (*domain.CSVRecord, error) {
	return s.repo.Update(id, record)
}

func (s *CSVRecordService) DeleteCSVRecord(id string) error {
	return s.repo.Delete(id)
}

func (s *CSVRecordService) GetAllCSVRecords(page int, limit int) ([]domain.CSVRecord, int64, error) {
	return s.repo.GetAll(page, limit)
}

func (s *CSVRecordService) processRecords(recordsChan <-chan domain.CSVRecord, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for record := range recordsChan {
		err := s.repo.Create(&record)
		if err != nil {
			log.Printf("Worker %d: Failed to insert record: %v", workerID, err)
		} else {
			log.Printf("Worker %d: Successfully inserted record: %v", workerID, record)
		}
	}
}
