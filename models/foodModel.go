package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required"`
	PRice      *float64           `json:"price" validate:"required"`
	Image_Url  *string            `json:"image_url" validate:"required"`
	Created_At time.Time          `json:"created_at" validate:"required"`
	Updated_At time.Time          `json:"updated_at" validate:"required"`
	Food_id    string             `json:"food_id" validate:"required"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}
