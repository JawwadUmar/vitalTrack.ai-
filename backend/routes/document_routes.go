package routes

import (
	"vita-track-ai/service"

	"github.com/gin-gonic/gin"
)

func createDocument(c *gin.Context) {
	service.CreateDocument(c)
}

func getDocument(c *gin.Context) {
	service.GetDocument(c)
}

func deleteDocument(c *gin.Context) {
	service.DeleteDocument(c)
}

func getCalendarDocuments(c *gin.Context) {
	service.GetCalendarDocuments(c)
}
