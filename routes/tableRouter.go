package routes

import "github.com/gin-gonic/gin"

func TableRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/table/:id", controller.GetTable())
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.PUT("/table/:id", controller.UpdateTable())
}
