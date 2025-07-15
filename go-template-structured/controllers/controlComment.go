package controllers

import (
	"context"
	"go-template/helpers"
	"go-template/models"
	"go-template/services"
	"net/http"
	"time"
	"strings"

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
	commentsCol := services.Client.Database("commentsdb").Collection("comments")
	usersCol := services.Client.Database("commentsdb").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convertir ID del artículo
	idStr := c.Param("id")
	articleID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	// Construir filtro base
	filter := bson.M{"article_id": articleID}

	// Filtro por fecha
	dateFilter := bson.M{}
	from := c.Query("from")
	to := c.Query("to")
	if from != "" {
		if fromTime, err := time.Parse("2006-01-02", from); err == nil {
			dateFilter["$gte"] = fromTime
		}
	}
	if to != "" {
		if toTime, err := time.Parse("2006-01-02", to); err == nil {
			toTime = toTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second + 999*time.Millisecond)
			dateFilter["$lte"] = toTime
		}
	}
	if len(dateFilter) > 0 {
		filter["created_at"] = dateFilter
	}

	// Obtener comentarios
	var comments []models.Comment
	cursor, err := commentsCol.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los comentarios"})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los comentarios"})
		return
	}

	// Obtener los IDs únicos de autores
	authorIDsSet := make(map[primitive.ObjectID]bool)
	for _, comment := range comments {
		authorIDsSet[comment.AuthorID] = true
	}
	var authorIDs []primitive.ObjectID
	for id := range authorIDsSet {
		authorIDs = append(authorIDs, id)
	}

	// Buscar los autores
	userFilter := bson.M{"_id": bson.M{"$in": authorIDs}}
	userCursor, err := usersCol.Find(ctx, userFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}
	defer userCursor.Close(ctx)

	type User struct {
		ID   primitive.ObjectID `bson:"_id"`
		Name string             `bson:"name"`
	}
	userMap := make(map[primitive.ObjectID]string)
	for userCursor.Next(ctx) {
		var user User
		if err := userCursor.Decode(&user); err == nil {
			userMap[user.ID] = user.Name
		}
	}

	// Combinar los comentarios con el nombre del autor
	type CommentWithAuthor struct {
		models.Comment `bson:",inline"`
		AuthorName     string `json:"author_name"`
	}

	var result []CommentWithAuthor
	for _, comment := range comments {
		name := userMap[comment.AuthorID]
		if searchAuthor := c.Query("author"); searchAuthor != "" {
			if !strings.Contains(strings.ToLower(name), strings.ToLower(searchAuthor)) {
				continue // filtrar por nombre de autor en Go
			}
		}
		result = append(result, CommentWithAuthor{
			Comment:     comment,
			AuthorName:  name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"comments": result})
}

