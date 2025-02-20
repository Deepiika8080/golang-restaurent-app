package routes

import (
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/food/:id", controller.GetFood())
	incomingRoutes.POST("/food", controller.CreateFood())
	incomingRoutes.PUT("/food/:id", controller.UpdateFood())

}
