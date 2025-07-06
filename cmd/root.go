package cmd

import (
	"os"
	"tourmate/feedback-service/constant/env"
	"tourmate/feedback-service/utils"

	"github.com/gin-gonic/gin"
)

// Execute server application
func Execute() {
	// Initialize logger config
	var logger = utils.GetLogConfig()

	// Load env
	loadEnv(logger)

	// Initialize gin server for API
	var server = gin.Default()

	// Config CORS for requests
	corsConfig(server)

	// Get API port
	var apiPort string = os.Getenv(env.API_PORT)

	// Set up API routes
	setupApiRoutes(server, apiPort)

	// Set up swagger
	setupSwagger(server, apiPort)

	// Run server
	if err := server.Run(":" + apiPort); err != nil {
		logger.Println("Error run server - " + err.Error())
	}
}
