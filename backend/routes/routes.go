package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api/v1")

	registerUserRoutes(api)
	registerFileRoutes(api)
}

func registerUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("/signup", signup)
		users.POST("/login", login)
		users.POST("/google", googleLogin)
	}
}

func registerFileRoutes(rg *gin.RouterGroup) {
	files := rg.Group("/files")
	{
		files.POST("/upload", uploadFile)
		files.GET("/:id", getFile)
	}
}
