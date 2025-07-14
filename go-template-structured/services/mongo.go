package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var db string = "commentsdb"

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Verificar la conexión
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB")
}

// GetCollection retorna una colección de MongoDB de forma segura
func GetCollection(collectionName string) (*mongo.Collection, error) {
	if Client == nil {
		return nil, fmt.Errorf("MongoDB client is not initialized")
	}
	return Client.Database(db).Collection(collectionName), nil
}

func OpenCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized. Please call InitMongo first.")
	}
	return Client.Database(db).Collection(collectionName)
}
