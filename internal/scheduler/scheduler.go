package scheduler

import (
	"fmt"
	"log"
	"os"
	"time"

	gocron "github.com/go-co-op/gocron/v2"
	logService "github.com/nielchaudhary/compass/pkg/logger"
)

func SchedulerCore() {

	logger := logService.GetLogger()

	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", err)
		}
	}()

	scheduler, newSchedulerError := gocron.NewScheduler()
	if newSchedulerError != nil {
		log.Fatal("Error Initialising the core scheduler", newSchedulerError)
	}

	repeatInterval := gocron.DurationJob(5 * time.Minute)
	healthCheckTask := gocron.NewTask(func() {
		_, err := PingEndpoint("http://localhost:3000/health", "GET")
		if err != nil {
			logger.Error("Error pinging endpoint")
		}
	})

	_, newJobError := scheduler.NewJob(repeatInterval, healthCheckTask)
	if newJobError != nil {
		logger.Error("Error running the cron job ", newJobError)
	}

	scheduler.Start()

}
