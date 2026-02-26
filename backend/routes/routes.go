package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api/v1")
	registerRoutesForUser(api)
	// registerRoutesForFiles(api)
}

func registerRoutesForUser(rg *gin.RouterGroup) {
	rg.POST("/signup", signup)
	rg.POST("/login", login)
	rg.POST("/google", googleLogin)
}
