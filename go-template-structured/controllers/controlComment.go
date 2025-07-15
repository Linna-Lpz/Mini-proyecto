package controllers

import (
	"context"
	"go-template/helpers"
	"go-template/models"
	"go-template/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	collection := services.Client.Database("commentsdb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Obtener article_id del path
	idStr := c.Param("id")
	articleID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	// Obtener parámetros de búsqueda
	searchAuthor := c.Query("author")
	from := c.Query("from")
	to := c.Query("to")

	matchStage := bson.M{
		"article_id": articleID,
	}

	// Agregar filtro de fechas si aplica
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
		matchStage["created_at"] = dateFilter
	}

	// Pipeline de agregación
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: matchStage}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "author_id",
				"foreignField": "_id",
				"as":           "author_info",
			},
		}},
		{{Key: "$unwind", Value: "$author_info"}},
	}

	// Agregar filtro por nombre del autor si viene
	if searchAuthor != "" {
		pipeline = append(pipeline, bson.D{{
			Key: "$match", Value: bson.M{
				"author_info.name": bson.M{"$regex": searchAuthor, "$options": "i"},
			},
		}})
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los comentarios"})
		return
	}
	defer cursor.Close(ctx)

	// Resultado extendido con nombre del autor
	type CommentWithAuthor struct {
		models.Comment `bson:",inline"`
		AuthorName     string `bson:"author_info.name" json:"author_name"`
	}

	var comments []CommentWithAuthor
	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los comentarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}
