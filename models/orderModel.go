package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id" validate:"required"`
	Order_Date time.Time          `json:"order_date" validate:"required"`
	Created_At *time.Time         `json:"created_at" validate:"required"`
	Updated_At *time.Time         `json:"updated_at" validate:"required"`
	Table_id   string             `json:"table_id" validate:"required"`
	Order_id   string             `json:"order_id" validate:"required"`
}
