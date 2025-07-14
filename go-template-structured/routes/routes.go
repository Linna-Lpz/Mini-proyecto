package routes

import (
    "github.com/gin-gonic/gin"
    "go-template/controllers"
    "go-template/middleware"
)

func CreateAuthentication(router *gin.Engine) {
    // Configurar rutas no protegidas
    router.GET("/articles", controllers.GetArticles)           // Obtener todos los artículos
    router.GET("/articles/:id", controllers.GetArticleById)    // Obtener un artículo por ID
    router.POST("/signup", controllers.Signup())                // Registrar un nuevo usuario
    router.POST("/login", controllers.Login())                 // Iniciar sesión

    protected := router.Group("/")
    protected.Use(middleware.Authenticate())
    {
        // Configurar rutas protegidas
        
    }
}

// CreateArticleRoutes: Configura las rutas para los artículos.
func CreateArticleRoutes(router *gin.Engine) {
    articleGroup := router.Group("/articles")
    {
        articleGroup.POST("/articles", controllers.CreateArticle)  // Crear un nuevo artículo
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