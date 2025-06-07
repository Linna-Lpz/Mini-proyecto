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

// Create: Crea un nuevo artículo.
func Create(c *gin.Context) {
	var article models.Article

	// Vincular el cuerpo de la solicitud al modelo Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	collection := services.Client.Database("commentsdb").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el artículo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Artículo creado con éxito",
		"_id":     res.InsertedID,
	})
}

// Get: Lista todos los artículos.
func Get(c *gin.Context) {
	collection := services.Client.Database("commentsdb").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los artículos"})
		return
	}
	defer cursor.Close(ctx)

	var articles []models.Article
	if err = cursor.All(ctx, &articles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los artículos"})
		return
	}

	c.JSON(http.StatusOK, articles)
}

// Get:  Devuelve un artículo con sus comentarios.
func GetArticle(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	collection := services.Client.Database("commentsdb").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var article models.Article
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&article)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	// Obtener comentarios asociados al artículo
	commentsCollection := services.Client.Database("commentsdb").Collection("comments")
	commentsCursor, err := commentsCollection.Find(ctx, bson.M{"articleid": idParam})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los comentarios"})
		return
	}
	defer commentsCursor.Close(ctx)

	var comments []models.Comment
	if err = commentsCursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los comentarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article":  article,
		"comments": comments,
	})
}

// Post: Crea un comentario para un artículo.
func PostComment(c *gin.Context) {
	idParam := c.Param("id")
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Asignar el ID del artículo al comentario (como string)
	comment.ArticleID = idParam

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
