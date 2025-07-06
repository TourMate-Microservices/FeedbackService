package cmd

import (
	"tourmate/feedback-service/utils"
)

// Execute server application
func Execute() {
	// Initialize logger config
	var logger = utils.GetLogConfig()

	// Load env
	loadEnv(logger)

	// Setup API routes
	setupApiRoutes(logger)

	// Setup gRPC routes
	setupGrpcRoutes(logger)
}
