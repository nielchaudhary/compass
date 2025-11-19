package storage

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	mongoClient *mongo.Client
	log         *zap.SugaredLogger
)

func init() {
	log = logger.GetLogger("internal/storage/mongoDB")
}

func ConnectMongoDB() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Warning: .env file not found or could not be loaded")
	}

	var mongoURI string
	if mongoURI = os.Getenv("mongoURI"); mongoURI == "" {
		log.Fatal("You must set your 'mongoURI' environment variable.")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(mongoURI).
		SetMaxPoolSize(50).
		SetMinPoolSize(10).
		SetServerAPIOptions(serverAPI)

	var err error
	mongoClient, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping to verify connection
	var result bson.M
	if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}
	log.Info("Pinged your deployment. You successfully connected to MongoDB!")
}

func GetDatabase(dbName string) *mongo.Database {
	if mongoClient == nil {
		log.Info("MongoDB client is not initialized. Please call ConnectMongoDB first.")

	}
	ConnectMongoDB()
	return mongoClient.Database(dbName)
}
