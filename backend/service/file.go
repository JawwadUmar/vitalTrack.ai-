package service

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"vita-track-ai/models"
	"vita-track-ai/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFiles(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "files are required",
		})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no files uploaded",
		})
		return
	}

	allowed := map[string]bool{
		".pdf": true,
		".jpg": true,
		".png": true,
	}

	var response []gin.H

	bucket := os.Getenv("AWS_BUCKET_NAME")

	for _, file := range files {

		// normalize extension
		ext := strings.ToLower(filepath.Ext(file.Filename))

		// validate extension
		if !allowed[ext] {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid file type: " + file.Filename,
			})
			return
		}

		// generate unique filename (S3 object key)
		storedName := uuid.New().String() + ext

		// upload file to S3
		if err := UploadToS3(file, storedName, bucket); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to upload to s3",
			})
			return
		}

		// create DB model
		fileModel := models.File{
			OriginalName: file.Filename,
			StoredName:   storedName,
			S3Key:        storedName, // ✔ correct S3 key
			FileSize:     file.Size,
			MimeType:     file.Header.Get("Content-Type"),
		}

		// save metadata to DB
		if err := repository.CreateFile(&fileModel); err != nil {

			// optional: rollback S3 upload if DB fails
			// _ = DeleteFromS3(bucket, storedName)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to save file metadata",
			})
			return
		}

		// response object
		response = append(response, gin.H{
			"file_id":       fileModel.ID,
			"original_name": fileModel.OriginalName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"files": response,
	})
}

func GetFileDownloadURL(fileID string) (string, error) {

	file, err := repository.GetFileByID(fileID)
	if err != nil {
		return "", err
	}

	bucket := os.Getenv("AWS_BUCKET_NAME")

	url, err := GenerateSignedURL(bucket, file.S3Key)
	if err != nil {
		return "", err
	}

	return url, nil
}
