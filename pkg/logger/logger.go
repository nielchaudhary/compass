package logger

import (
	"log"

	constants "github.com/nielchaudhary/compass/pkg/constants"
	"go.uber.org/zap"
)

var (
	// globalLogger is the global logger instance used throughout the application
	globalLogger *zap.Logger
	// sugarLogger is a sugared logger for easier usage
	sugarLogger *zap.SugaredLogger
)

// InitLogger initializes the global zap logger instance based on the environment.
// It accepts the environment (env) as a parameter and configures the logger accordingly.
func InitLogger(env string) error {
	var logger *zap.Logger
	var err error

	switch env {
	case string(constants.Development):
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Failed to initialize development logger: %v", err)
			return err
		}
	case string(constants.Production):
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to initialize production logger: %v", err)
			return err
		}
	case string(constants.Testing):
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Failed to initialize testing logger: %v", err)
			return err
		}
	default:
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Failed to initialize default logger: %v", err)
			return err
		}
	}

	globalLogger = logger
	sugarLogger = logger.Sugar()

	return nil
}

func GetLogger(fileName string) *zap.SugaredLogger {
	if sugarLogger == nil {
		// If logger is not initialized, initialize with development settings
		_ = InitLogger(string(constants.Development))
	}

	// Return a logger with the fileName field added
	return sugarLogger.With("file", fileName)
}

// Sync flushes any buffered log entries.
// Applications should call Sync before exiting.
func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}
