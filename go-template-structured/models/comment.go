package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Comment: Comentario de un usuario acerca de un artículo
type Comment struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Author      string              `json:"author" bson:"author"`
    Text        string              `json:"text" bson:"text"`
    CreatedAt   string              `json:"created_at" bson:"created_at"`
    ArticleID   string              `json:"article_id" bson:"article_id"`
}