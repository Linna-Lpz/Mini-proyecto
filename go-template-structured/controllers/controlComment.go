package controllers

import (
    "github.com/gin-gonic/gin"
    "go-template/models"
    "go-template/services"
	"net/http"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

// GetComment: Obtiene una lista de comentarios para un artículo específico
func GetComments(c *gin.Context){
	collection := services.Client.Database("commentsdb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	articleID := c.Param("id")
	filter := bson.M{"article_id": articleID}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los comentarios"})
		return
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los comentarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})


}