package core

import (
	"context"
	"time"

	"github.com/nielchaudhary/compass/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// will try to maximise its utility across the codebase, hence using any as data type
func BatchData[T any](data []T, size int) []T {
	batch := make([]T, size)
	for i := 0; i < len(data) && i < size; i++ {
		batch[i] = data[i]
	}

	return batch
}

// contains all the db related functions that would be used in core engine
func FetchEndpointsDataFromDB(collectionName string) ([]Endpoints, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err := storage.GetCollection(collectionName)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []Endpoints
	for cursor.Next(context.TODO()) {
		var result Endpoints
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
