package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date string             `json:"start_date" validate:"required"`
	End_Date   string             `json:"end_date" validate:"required"`
	Created_At time.Time          `json:"created_at" validate:"required"`
	Updated_At time.Time          `json:"updated_at" validate:"required"`
	menu_id    string             `json:"food_id" validate:"required"`
}
