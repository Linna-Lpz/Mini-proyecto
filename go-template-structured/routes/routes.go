package routes

import (
    "github.com/gin-gonic/gin"
    "go-template/controllers"
)

func RegisterRoutes(router *gin.Engine) {
    router.GET("/ping", controllers.PingController)
}

func CreateArticleRoutes(router *gin.Engine) {
    articleGroup := router.Group("/articles")
    {
        articleGroup.POST("/", controllers.Create)  // Crrear un nuevo artículo
        articleGroup.GET("/", controllers.Get)  // Obtener todos los artículos
        articleGroup.GET("/:id", controllers.GetArticle)    // Obtener un artículo por ID
        articleGroup.POST("/:id/comments", controllers.PostComment) // Publicar un comentario en un artículo
    }
}