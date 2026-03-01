package routes

import (
	"net/http"
	"vita-track-ai/service"

	"github.com/gin-gonic/gin"
)

func uploadFile(c *gin.Context) {
	service.UploadFiles(c)
}

func getFile(c *gin.Context) {

	id := c.Param("id")

	url, err := service.GetFileDownloadURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate download url",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
