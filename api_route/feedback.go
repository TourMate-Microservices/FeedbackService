package apiroute

import "github.com/gin-gonic/gin"

func InitializeFeedbackHandlerRoute(server *gin.Engine, port string) {
	// Context path
	//var contextPath string = "Feedbacks"

	// Define Feedback endpoints with admin required
	//var adminAuthGroup = server.Group(contextPath, middleware.Authorize, middleware.AdminAuthorization)
	// adminAuthGroup.GET("", handler.GetAllFeedbacks)
	// adminAuthGroup.PUT("/update", handler.UpdateFeedback)

	// // Define Feedback endpoints with basic required
	// var authGroup = server.Group(contextPath, middleware.Authorize)
	// authGroup.GET("/user/:id", handler.GetFeedbacksByUser)
	// authGroup.GET("/:id", handler.GetFeedbackById)
}
