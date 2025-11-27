package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nielchaudhary/compass/internal/config"
	constants "github.com/nielchaudhary/compass/pkg/constants"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	zap "go.uber.org/zap"

	storage "github.com/nielchaudhary/compass/internal/storage"
)

var zapLog *zap.SugaredLogger

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	coreServer := fiber.New()
	env := config.GetEnv("APP_ENV", string(constants.Development))
	serverPort := config.GetEnv("PORT", "8090")

	storage.ConnectMongoDB()

	if err := logger.InitLogger(env); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", err)
		}
	}()

	zapLog = logger.GetLogger()

	zapLog.Infow("STARTING SERVER",
		"server", "compass-core",
		"environment", env,
		"port", serverPort,
	)

	if err := coreServer.Listen(":" + serverPort); err != nil {
		zapLog.Fatalw("CORE SERVER BOOTUP FAILED", "error", err)
		os.Exit(1)

	}

}
