package router

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	monitor "github.com/nielchaudhary/compass/internal/monitor/api"
	logger "github.com/nielchaudhary/compass/pkg/logger"
)

func InitMonitorRouter() {
	log := logger.GetLogger()
	if log == nil {
		log.Panic("COMPASS LOGGER CRASHED: FAILED TO INITIALIZE LOGGER DUE TO: logger is nil")
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", err)
		}
	}()

	//router
	monitorRouter := fiber.New(fiber.Config{
		ServerHeader: "compass-core/monitor",
	})
	monitorRouter.Get("/compass/v1/register-endpoint", monitor.RegisterServiceEndpoints)

	monitorServerPort := ":8080"
	log.Infow("COMPASS MONITOR API SERVER V1 STARTING ON PORT 8080")
	err := monitorRouter.Listen(monitorServerPort)
	if err != nil {
		log.Fatalw("COMPASS MONITOR API SERVER V1 CRASHED: FAILED TO START SERVER", "error", err)
	}

	log.Info("Initialised Monitor Router")
}
