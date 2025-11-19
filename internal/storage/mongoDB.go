package storage

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() {

	log := logger.GetLogger("internal/storage/mongoDB")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Warning: .env file not found or could not be loaded")
	}

	var mongoURI string

	if mongoURI = os.Getenv("mongoURI"); mongoURI == "" {
		log.Fatal("You must set your 'mongoURI' environment variable.")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	log.Info("Pinged your deployment. You successfully connected to MongoDB!")

}
