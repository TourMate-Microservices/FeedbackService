package apiroute

import (
	"github.com/gin-gonic/gin"
)

func InitializePlatformFeedbackHandlerRoute(server *gin.Engine, port string) {
	// Context path
	//var contextPath string = "PlatformFeedbacks"

	// Define PlatformFeedback endpoints with admin required
	//var adminAuthGroup = server.Group(contextPath, middleware.Authorize, middleware.AdminAuthorization)
	// adminAuthGroup.GET("", handler.GetAllPlatformFeedbacks)
	// adminAuthGroup.PUT("/update", handler.UpdatePlatformFeedback)

	// // Define PlatformFeedback endpoints with basic required
	// var authGroup = server.Group(contextPath, middleware.Authorize)
	// authGroup.GET("/user/:id", handler.GetPlatformFeedbacksByUser)
	// authGroup.GET("/:id", handler.GetPlatformFeedbackById)
}
