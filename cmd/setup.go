package cmd

import (
	"log"
	"os"
	"tourmate/feedback-service/constant/env"
	"tourmate/feedback-service/docs"
	api "tourmate/feedback-service/route/api"

	_ "tourmate/feedback-service/docs"

	"github.com/gin-gonic/gin"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
)

func setupApiRoutes(logger *log.Logger) {
	// Initialize gin server for API
	var server = gin.Default()

	// Config CORS for requests
	corsConfig(server)

	// Get API port
	var apiPort string = os.Getenv(env.API_PORT)

	// Set up swagger
	setupSwagger(server, apiPort)

	// Feedback API endpoints
	api.InitializeFeedbackHandlerRoute(server, apiPort)

	// Platform Feedback API endpoints
	api.InitializePlatformFeedbackHandlerRoute(server, apiPort)

	// Run server
	if err := server.Run(":" + apiPort); err != nil {
		logger.Println("Error run service - " + err.Error())
	}
}

func setupSwagger(server *gin.Engine, port string) {
	//Configure swagger info
	docs.SwaggerInfo.Title = "Tourmate - Feedback Service API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:" + port

	//Add swagger route
	server.GET("swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler))
}

func setupGrpcRoutes(logger *log.Logger) {

}
