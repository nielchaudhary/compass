package constants

type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
	Testing     Environment = "Testing"
)

type validMethods []string

var ValidHttpMethods = validMethods{"GET", "POST", "PUT", "PATCH", "DELETE"}
