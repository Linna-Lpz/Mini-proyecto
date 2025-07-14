package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User: Representa un usuario en el sistema
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" validate:"required,min=2" bson:"name"`
	Email    string             `json:"email" validate:"email,required" bson:"email"`
	Password string             `json:"password" validate:"required,min=6" bson:"password"`
	Token         *string            `json:"token" bson:"token,omitempty"`
	Refresh_token *string            `json:"refresh_token" bson:"refresh_token,omitempty"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}