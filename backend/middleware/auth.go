package middleware

import (
	"net/http"
	"strings"
	"vita-track-ai/utility"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		// Expect: Bearer TOKEN
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if err := utility.VerifyToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		userId, err := utility.GetUserIdFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		// ⭐ store user id in context
		c.Set("user_id", userId)

		c.Next()
	}
}
