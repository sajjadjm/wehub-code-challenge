package ports

import (
	"github.com/sajjadjm/wehub-code-challenge/internal/core/domain"
)

// CSVReaderInterface defines the interface for reading CSV data.
type CSVReaderInterface interface {
	ReadCSV(filePath string) ([]domain.CSVRecord, error)
}

type CSVRecordRepositoryInterface interface {
	Create(record *domain.CSVRecord) error
	BulkCreate(records []domain.CSVRecord) error
	GetByID(id string) (*domain.CSVRecord, error)
	Update(id string, record *domain.CSVRecord) (*domain.CSVRecord, error)
	Delete(id string) error
	GetAll(page int, limit int) ([]domain.CSVRecord, int64, error)
}
