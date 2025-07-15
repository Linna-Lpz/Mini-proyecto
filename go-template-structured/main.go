package main

import (
	"go-template/helpers"
	"go-template/routes"
	"go-template/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    // Inicializar conexión Mongo
    services.InitMongo()

    // Key
    helpers.SetJWTKey("key")


    r := gin.Default()
    r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))


    // Registrar rutas
    routes.CreateArticleRoutes(r)
    routes.CreateAuthentication(r)

    // Iniciar el servidor
    r.Run(":8080")
}