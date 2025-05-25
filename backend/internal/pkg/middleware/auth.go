package middleware

import (
	"net/http"
	
	"github.com/BookIT/backend/internal/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserKey = "X-Auth-Token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(ContextUserKey)
		if tokenString == "" {
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
			
			c.AbortWithStatusJSON(status, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}