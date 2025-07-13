package controllers

import (
	"context"
	"go-template/models"
	"go-template/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post: Crea un comentario para un artículo.
func PostComment(c *gin.Context) {
	idParam := c.Param("id")
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Convertir idParam a ObjectID y asignar al comentario
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}
	comment.ArticleID = objID

	commentsCollection := services.Client.Database("commentsdb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := commentsCollection.InsertOne(ctx, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el comentario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comentario creado con éxito",
		"_id":     res.InsertedID,
	})
}

// GetComments: Obtiene los comentarios de un artículo específico, aplicando un filtro.
func GetComments(c *gin.Context) {
	collection := services.Client.Database("commentsdb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtros
	filter := bson.M{}

	// Filtro por article_id
	articleIDParam := c.Param("id")
	if articleIDParam != "" {
		if articleObjID, err := primitive.ObjectIDFromHex(articleIDParam); err == nil {
			filter["article_id"] = articleObjID
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
			return
		}
	}

	// Filtro por autor (author_id)
	author := c.Query("author")
	if author != "" {
		if authorObjID, err := primitive.ObjectIDFromHex(author); err == nil {
			filter["author_id"] = authorObjID
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de autor inválido"})
			return
		}
	}

	// Filtro por rango de fechas
	from := c.Query("from")
	to := c.Query("to")
	dateFilter := bson.M{}
	if from != "" {
		if fromTime, err := time.Parse("2006-01-02", from); err == nil {
			dateFilter["$gte"] = fromTime
		}
	}
	if to != "" {
		if toTime, err := time.Parse("2006-01-02", to); err == nil {
			dateFilter["$lte"] = toTime
		}
	}
	if len(dateFilter) > 0 {
		filter["created_at"] = dateFilter
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los artículos"})
		return
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los artículos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}
