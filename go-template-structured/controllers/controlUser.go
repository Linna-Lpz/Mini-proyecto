package controllers

import (
	"context"
	"go-template/helpers"
	"go-template/models"
	"go-template/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var userCollection *mongo.Collection

func getUserCollection() (*mongo.Collection, error) {
	var err error
	if userCollection == nil {
		userCollection, err = services.GetCollection("users")
		if err != nil {
			return nil, err
		}
	}
	return userCollection, nil
}

// GetUser: Obtiene un usuario por ID.
func GetUser(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	collection, err := getUserCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to database"})
		return
	}

	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Signup: Registra un nuevo usuario.
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		//Get user input
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Validate user input
		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		collection, err := getUserCollection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to database"})
			return
		}

		count, err := collection.CountDocuments(ctx, bson.M{
			"$or": []bson.M{
				{"email": user.Email},
				{"name": user.Name},
				{"userId": user.ID.Hex()},
			},
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}

		//Generate rest of the user data
		hashedPwd := helpers.HashPassword(&user.Password)
		user.Password = *hashedPwd
		user.ID = primitive.NewObjectID()
		accessToken, refreshToken := helpers.GenerateTokens(user.Email, user.Name, user.ID.Hex())
		user.Token = &accessToken
		user.Refresh_token = &refreshToken

		_, insertErr := collection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

// Login: Inicia sesión de un usuario.
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		// Get user input
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection, err := getUserCollection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to database"})
			return
		}

		err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}

		//verify password
		passwordIsValid, msg := helpers.VerifyPassword(foundUser.Password, user.Password)
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		token, refreshToken := helpers.GenerateTokens(foundUser.Email, foundUser.Name, foundUser.ID.Hex())
		helpers.UpdateAllTokens(token, refreshToken, foundUser.ID.Hex())

		c.JSON(http.StatusOK, gin.H{
			"user":          foundUser,
			"token":         token,
			"refresh_token": refreshToken,
		})
	}
}
