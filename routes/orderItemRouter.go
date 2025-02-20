package routes

import (
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItem/:id", controller.GetOrderItem())
	incomingRoutes.POST("/orderItem", controller.createOrderItem())
	incomingRoutes.PUT("/orderItem", controller.updateOrderItem())
}
