package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id" validate:"required"`
	Number_Of_guests *int               `json:"number_of_guests" validate:"required"`
	Table_no         *int               `json:"table_no" validate:"required"`
	Created_At       *time.Time         `json:"created_at" validate:"required"`
	Updated_At       *time.Time         `json:"updated_at" validate:"required"`
	Table_id         string             `json:"table_id" validate:"required"`
}
