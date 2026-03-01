package main

import (
	"vita-track-ai/config"
	"vita-track-ai/database"
	"vita-track-ai/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Init()
	config.InitS3()
	var server *gin.Engine = gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8081")

}
