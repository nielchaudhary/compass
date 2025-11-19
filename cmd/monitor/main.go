package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nielchaudhary/compass/internal/config"
	monitor "github.com/nielchaudhary/compass/internal/monitor/router"
	storage "github.com/nielchaudhary/compass/internal/storage"
	constants "github.com/nielchaudhary/compass/pkg/constants"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	zap "go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {

	storage.ConnectMongoDB()

	monitorServer := fiber.New()

	monitor.InitMonitorRouter()

	env := config.GetEnv("APP_ENV", string(constants.Development))
	monitorPort := config.GetEnv("MONITOR_PORT", "8080")

	if err := logger.InitLogger(env); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", err)
		}
	}()

	log = logger.GetLogger("cmd/monitor/main.go")

	log.Infow("STARTING SERVER",
		"server", "compass-monitor",
		"environment", env,
		"port", monitorPort,
	)

	if err := monitorServer.Listen(":" + monitorPort); err != nil {
		log.Fatalw("Server failed", "error", err)
		os.Exit(1)
	}

}
