package logger

import (
	"fmt"
)

// InitLogger initializes and returns a zap logger instance.
// It accepts the environment (env) as a parameter and returns a Logger.
func InitLogger(env string, fileName string) string {

	fmt.Println("Initialising Logger with environemnt : ", env)
	return "test"

}
