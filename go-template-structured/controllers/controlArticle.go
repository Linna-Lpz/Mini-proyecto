package controllers

import (
	"context"
	"go-template/models"
	"go-template/services"
	"go-template/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateArticle: Crea un nuevo artículo.
func CreateArticle(c *gin.Context) {
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

// GetArticles: Lista todos los artículos según filtro, con paginación.
func GetArticles(c *gin.Context) {
	collection := services.Client.Database("commentsdb").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtros
	filter := utils.FilterSearch(c, "title", "created_at")
	
	// Paginación
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	// Contar el total de artículos con filtro
	totalCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar los artículos"})
		return
	}

	// Opciones de paginación
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	cursor, err := collection.Find(ctx, filter, findOptions)
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

	c.JSON(http.StatusOK, gin.H{
		"data":  articles,
		"page":  page,
		"limit": limit,
		"total": totalCount,
	})
}

// Función auxiliar para obtener comentarios por ArticleID
func FetchCommentsByArticleID(ctx context.Context, articleID string) ([]models.Comment, error) {
	commentsCollection := services.Client.Database("commentsdb").Collection("comments")
	objtID, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"article_id": objtID}
	cursor, err := commentsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// GetArticleById:  Devuelve un artículo con sus comentarios.
func GetArticleById(c *gin.Context) {
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

	comments, err := FetchCommentsByArticleID(ctx, idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los comentarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article":  article,
		"comments": comments,
	})
}
