package cmd

import (
	api_route "tourmate/feedback-service/api_route"
	//_ "tourmate/feedback-service/docs"

	"github.com/gin-gonic/gin"
	//swagger_files "github.com/swaggo/files"
	//gin_swagger "github.com/swaggo/gin-swagger"
)

func setupApiRoutes(server *gin.Engine, port string) {
	// Feedback API endpoints
	api_route.InitializeFeedbackHandlerRoute(server, port)

	// Platform Feedback API endpoints
	api_route.InitializePlatformFeedbackHandlerRoute(server, port)
}

func setupSwagger(server *gin.Engine, port string) {
	// Configure swagger info
	// docs.SwaggerInfo.Title = "Sonit Server API"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// docs.SwaggerInfo.Host = "localhost:" + port

	// Add swagger route
	//server.GET("swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler))
}
