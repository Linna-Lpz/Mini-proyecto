package main

import (
	"go-template/config"
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
    Key := config.GenerateRandomKey()
    helpers.SetJWTKey(Key)

    r := gin.Default()
    r.Use(cors.Default())

    // Registrar rutas
    routes.CreateArticleRoutes(r)
    routes.CreateCommentRoutes(r)
    routes.CreateAuthentication(r)

    // Iniciar el servidor
    r.Run(":8080")
}