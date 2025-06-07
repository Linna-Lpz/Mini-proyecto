package controllers

import (
    "github.com/gin-gonic/gin"
    "go-template/models"
    "go-template/services"
	"context"
	"time"
)

// CreateComment: Crea un nuevo comentario para un artículo.
func CreateComment(c *gin.Context){
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	collection := services.Client.Database("commentsdb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validar el cuerpo de la solicitud
	res, err := collection.InsertOne(ctx, comment)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al crear el comentario"})
		return
	}

	// Respuesta exitosa
	c.JSON(201, gin.H{
		"message": "Comentario creado con éxito",
		"id": res.InsertedID,
	})
}