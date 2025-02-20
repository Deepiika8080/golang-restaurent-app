package routes

import (
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:id", controller.GetInvoice())
	incomingRoutes.POST("/invoice", controller.createInvoice())
	incomingRoutes.PUT("/invoice", controller.updateInvoice())
}
