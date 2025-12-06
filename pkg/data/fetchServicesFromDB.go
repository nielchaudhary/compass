package data

import (
	db "github.com/nielchaudhary/compass/internal/storage"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchServiceEndpointsFromDB() (*mongo.Collection, error) {

	log := logger.GetLogger()

	endpointsColl, err := db.GetCollection("endpoints")

	if err != nil {
		log.Error("Error fetching endpointsColl from DB due to : ", err)
		return nil, err
	}

	return endpointsColl, nil

}
