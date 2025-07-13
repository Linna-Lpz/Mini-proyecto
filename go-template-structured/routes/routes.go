package routes

import (
    "github.com/gin-gonic/gin"
    "go-template/controllers"
)

// CreateArticleRoutes: Configura las rutas para los artículos.
func CreateArticleRoutes(router *gin.Engine) {
    articleGroup := router.Group("/articles")
    {
        articleGroup.POST("/articles", controllers.CreateArticle)  // Crear un nuevo artículo
        articleGroup.GET("/articles", controllers.GetArticles)  // Obtener todos los artículos
        articleGroup.GET("/articles/:id", controllers.GetArticleById)    // Obtener un artículo por ID
    }
}

// CreateCommentRoutes: Configura las rutas para los comentarios.
func CreateCommentRoutes(router *gin.Engine) {
    commentGroup := router.Group("/comments")
    {
        commentGroup.GET("/articles/:id/comments", controllers.GetComments) // Obtener comentarios de un artículo
        commentGroup.POST("/articles/:id/comments", controllers.PostComment) // Publicar un comentario en un artículo
    }
}

// CreateAuthentication: Configura las rutas de autenticación.
func CreateAuthentication(router *gin.Engine) {
    authGroup := router.Group("/auth")
    {
        authGroup.POST("/auth/register", controllers.Register) // Registrar un nuevo usuario
        authGroup.POST("/auth/login", controllers.Login) // Iniciar sesión
        authGroup.GET("/auth/me", controllers.GetProfile) // Obtener perfil del usuario
    }
}