package middleware

import (
	"net/http"
	"strings"
	"taskManager/api/helper-api/security"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware Context Authn
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorization"})
			ctx.Abort()
			return
		}

		// Split by space because in authorization is Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := security.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)

		ctx.Next()
	}
}
