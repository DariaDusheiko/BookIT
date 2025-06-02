package middleware

import (
	"net/http"

	"github.com/BookIT/backend/internal/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserKey = "X-Auth-Token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Всегда разрешаем OPTIONS запросы без проверки авторизации
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // Правильный статус для preflight
			c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
			c.Header("Access-Control-Allow-Credentials", "true")
			return
		}

		tokenString := c.GetHeader(ContextUserKey)
		if tokenString == "" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization token required",
			})
			return
		}

		userID, err := utils.ParseAndValidateToken(tokenString)
		if err != nil {
			status := http.StatusUnauthorized
			if err == utils.ErrExpiredToken {
				status = http.StatusForbidden
			}

			c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
			c.Header("Access-Control-Allow-Credentials", "true")

			c.AbortWithStatusJSON(status, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

// CorsMiddleware возвращает middleware с настройками CORS
func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", ContextUserKey},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
