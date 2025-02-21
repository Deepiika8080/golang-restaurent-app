package routes

import (
	controller "golang-restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItem/:id", controller.GetOrderItem())
	incomingRoutes.POST("/orderItem", controller.CreateOrderItem())
	incomingRoutes.PUT("/orderItem", controller.UpdateOrderItem())
}
