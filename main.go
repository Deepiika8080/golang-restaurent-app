package main

import (
	"golang-restaurent-management/database"
	"golang-restaurent-management/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.FoodRoutes(router)
	router.MenuRouter(router)
	router.OrderRoutes(router)
	router.TableRoutes(router)
	router.InvoiceRoutes(router)
	router.OrderItemRoutes(router)

	router.Run(":" + port)
}
