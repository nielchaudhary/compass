package config

import (
	"os"
)

// GetEnv retrieves an environment variable value or returns a fallback if not set.
//
// Parameters:
//   - key: environment variable name
//   - fallback: default value if key is not set or empty
//
// Returns: environment variable value or fallback

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value

}
