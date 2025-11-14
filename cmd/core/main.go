package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nielchaudhary/compass/internal/config"
	constants "github.com/nielchaudhary/compass/pkg/constants"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	zap "go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {
	env := config.GetEnv("APP_ENV", string(constants.Development))
	serverPort := config.GetEnv("PORT", "8090")

	if err := logger.InitLogger(env); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", err)
		}
	}()

	log = logger.GetLogger("cmd/core/main.go")

	log.Infow("STARTING SERVER",
		"server", "compass-core",
		"environment", env,
		"port", serverPort,
	)

	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Fatalw("Server failed", "error", err)
		os.Exit(1)

	}

}
