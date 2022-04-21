package routes

import (
	controller "restaurent-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())

}
