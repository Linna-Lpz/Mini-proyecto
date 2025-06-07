package main

import (
    "go-template/routes"
    "go-template/services"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    // Inicializar conexión Mongo
    services.InitMongo()

    r := gin.Default()
    r.Use(cors.Default())

    // Registrar rutas
    routes.RegisterRoutes(r)
    routes.CreateArticleRoutes(r)

    r.Run(":8080")
}