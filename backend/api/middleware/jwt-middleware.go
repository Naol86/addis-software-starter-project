package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naol86/addis-software-starter/project/backend/package/tokens"
)

func JwtAuthMiddleWare(secret string) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Bearer token is missing"})
			c.Abort()
			return
		}

		ok, err := tokens.VerifyToken(tokenString, secret)
		if err != nil || !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next();

	}
}