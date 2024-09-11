package db

import (
	"context"
	"fmt"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/domain"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CSVRecordRepository struct {
	collection *mongo.Collection
}

// NewCSVRecordRepository creates a new MongoRepository
func NewCSVRecordRepository(uri, dbName, collectionName string) (ports.CSVRecordRepositoryInterface, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &CSVRecordRepository{collection: collection}, nil
}

// Create inserts a single CSVRecord into MongoDB.
func (r *CSVRecordRepository) Create(record *domain.CSVRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		log.Fatal("Cannot insert record: ", err)
	}
	return nil
}

// BulkCreate inserts multiple CSVRecords into MongoDB.
func (r *CSVRecordRepository) BulkCreate(records []domain.CSVRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var docs []interface{}
	for _, record := range records {
		docs = append(docs, record)
	}

	_, _ = r.collection.InsertMany(ctx, docs)
	return nil
}

func (r *CSVRecordRepository) GetByID(id string) (*domain.CSVRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var record domain.CSVRecord
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&record)
	return &record, err
}

func (r *CSVRecordRepository) Update(id string, record *domain.CSVRecord) (*domain.CSVRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": record})

	var updatedRecord domain.CSVRecord
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedRecord)
	if err != nil {
		return nil, err
	}

	return &updatedRecord, nil
}

func (r *CSVRecordRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return nil
}

func (r *CSVRecordRepository) GetAll(page int, limit int) ([]domain.CSVRecord, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println("Error in closing mongo connection")
		}
	}(cursor, ctx)

	var records []domain.CSVRecord
	if err = cursor.All(ctx, &records); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}
