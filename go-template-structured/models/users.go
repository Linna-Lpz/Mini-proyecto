package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User: Representa un usuario en el sistema
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}