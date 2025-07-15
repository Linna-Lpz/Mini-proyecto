package controllers

import (
	"context"
	"go-template/utils"
	"go-template/helpers"
	"go-template/models"
	"go-template/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

	// Obtener el ID del autor desde el token JWT
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
		return
	}
	claims := claimsRaw.(*helpers.Claims)

	authorID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de autor inválido"})
		return
	}
	comment.AuthorID = authorID
	comment.CreatedAt = time.Now()

	// Insertar el comentario en la colección
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
	// Obtener la colección de comentarios
	collection := services.Client.Database("commentsdb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtros
	idStr := c.Param("id")
	articleID, err:= primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	filter := utils.FilterSearch(c, "author", "created_at")
	filter["article_id"] = articleID

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
