package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id" validate:"required"`
	Text       string             `json:"text" validate:"required"`
	Title      string             `json:"title" validate:"required"`
	Note_id    string             `json:"note_id" validate:"required"`
	Created_At *time.Time         `json:"created_at" validate:"required"`
	Updated_At *time.Time         `json:"updated_at" validate:"required"`
}
