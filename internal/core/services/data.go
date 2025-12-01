package core

import (
	"context"

	storage "github.com/nielchaudhary/compass/internal/storage"
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

func FetchDataFromDB(collectionName string, filter bson.M) ([]bson.M, error) {
	collection, err := storage.GetCollection(collectionName)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	for cursor.Next(context.TODO()) {
		var result bson.M
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
