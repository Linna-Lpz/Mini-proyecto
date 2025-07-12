package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article: Representa un artículo en el sistema
type Article struct {
    ID    		primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
    Title 		string 					`json:"title" bson:"title"`
	Content 	string 					`json:"content" bson:"content"`
	CreatedAt 	time.Time 				`json:"created_at" bson:"created_at"`
}