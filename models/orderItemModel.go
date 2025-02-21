package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id" validate:"required"`
	Unit_Price    *float64           `json:"unit_price" validate:"required"`
	Quantity      *string            `json:"string" validate:"required"`
	Created_At    *time.Time         `json:"created_at" validate:"required"`
	Updated_At    *time.Time         `json:"updated_at" validate:"required"`
	Order_id      string             `json:"order_id" validate:"required"`
	Food_id       *string            `json:"food_id" validate:"required"`
	Order_item_id string             `json:"order_item_id" validate:"required"`
}
