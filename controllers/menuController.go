package controller

import (
	"context"
	"fmt"
	"golang-restaurent-management/database"
	model "golang-restaurent-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.DBInstance(), "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		result, err := menuCollection.Find(context.Background(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu"})
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)
		defer cancel()

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		defer cancel()

		menuId := c.Param("id")
		var menu model.Menu
		err := menuCollection.FindOne(ctx, bson.H{"menu_id", menuId}).Decode(&menu)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food"})
		}
		c.JSON(http.StatusOK, menu)
		defer cancel()

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		defer cancel()

		var menu model.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		menu.Created_At,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_At,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		menu.ID = primitive.NewObjectID()
		result, insertErr := menuCollection.InsertOne(ctx, menu)

		if insertErr != nil {
			msg := fmt.Sprint("menu was not found!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func inTimeSpan(start,end,check time.Time) bool {
       return start.After(time.Now()) && end.After(start)
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
        var ctx,cancel = context.WithTimeout(context.TODO(),100*time.Second)
		defer cancel()

		var menu model.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		menuId := c.Param("id")
		filter := bson.M{
			"menu_id": menuId
		}

		var updateobj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date,*menu.End_Date, time.Now()) {
				msg := "kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return
			}

			updateobj = append(updateobj,bson.E{Key: "start_date", Value: menu.Start_Date})
			updateobj = append(updateobj,bson.E{Key: "end_date", Value: menu.End_Date})

			if menu.Category != "" {
				updateobj = append(updateobj,bson.E{Key: "category", Value: menu.Category})
			}

			if menu.Name != "" {
				updateobj = append(updateobj,bson.E{Key: "name", Value: menu.Category})
			}

			menu.Updated_At,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
			updateobj = append(updateobj,bson.E{Key: "updated_at", Value: menu.Updated_At})

			upsert := true

			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result,err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{

				}
			)
		}


	}
}
