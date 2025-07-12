package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment: Comentario de un usuario acerca de un artículo
type Comment struct {
    ID          primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
    ArticleID   primitive.ObjectID      `json:"article_id" bson:"article_id"`
    AuthorID    primitive.ObjectID      `json:"author_id" bson:"author_id"`
    Text        string                  `json:"text" bson:"text"`
    CreatedAt   time.Time               `json:"created_at" bson:"created_at"`
}