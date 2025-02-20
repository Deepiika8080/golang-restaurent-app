package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:id", controller.GetUser())
	incomingRoutes.POST("/user/signup", controller.SignUp())
	incomingRoutes.POST("/user/login", controller.Login())
}
