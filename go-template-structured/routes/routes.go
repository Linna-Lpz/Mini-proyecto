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
        articleGroup.POST("/", controllers.Create)
        articleGroup.GET("/", controllers.Get)
        articleGroup.GET("/:id", controllers.GetArticle)
        articleGroup.POST("/:id/comments", controllers.Create)
    }
}