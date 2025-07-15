package routes

import (
    "github.com/gin-gonic/gin"
    "go-template/controllers"
    "go-template/middleware"
)

func CreateAuthentication(router *gin.Engine) {
    // Configurar rutas no protegidas
    router.GET("/articles", controllers.GetArticles)                // Obtener todos los artículos
    router.GET("/articles/:id", controllers.GetArticleById)         // Obtener un artículo por ID
    router.GET("/articles/:id/comments", controllers.GetComments)   // Obtener comentarios de un artículo
    router.POST("/signup", controllers.Signup())                    // Registrar un nuevo usuario
    router.POST("/login", controllers.Login())                      // Iniciar sesión

    routerSecure := router.Group("/")
    routerSecure.Use(middleware.Authenticate())
    {
        // Configurar rutas protegidas
        routerSecure.POST("/articles/:id/comments", controllers.PostComment) // Publicar un comentario en un artículo
    }
}

// CreateArticleRoutes: Configura las rutas para los artículos.
func CreateArticleRoutes(router *gin.Engine) {
    articleGroup := router.Group("/articles")
    {
        articleGroup.POST("/articles", controllers.CreateArticle)  // Crear un nuevo artículo
    }
}