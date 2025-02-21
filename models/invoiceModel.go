package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string
	Order_id         *string
	Payment_Method   string
	Payment_status   string
	Payment_due_date time.Time
	Created_At       time.Time `json:"created_at" validate:"required"`
	Updated_At       time.Time `json:"updated_at" validate:"required"`
}
