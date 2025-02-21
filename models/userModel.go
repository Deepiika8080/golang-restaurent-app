package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" validate:"required"`
	Name          *string            `json:"name" validate:"required"`
	Password      *string            `json:"password" validate:"required"`
	Email         *string            `json:"email" validate:"required"`
	Avatar        *string            `json:"avatar" validate:"required"`
	Token         *string            `json:"token" validate:"required"`
	Refresh_Token *string            `json:"refresh_token" validate:"required"`
	User_id       *string            `json:"user_id" validate:"required"`
	Created_At    *time.Time         `json:"created_at" validate:"required"`
	Updated_At    *time.Time         `json:"updated_at" validate:"required"`
}
