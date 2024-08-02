package middlewares

import (
	"crud-go/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization");
		if authHeader == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            ctx.Abort()
            return
        }
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
            ctx.Abort()
            return
        }
		claims, err := jwt.ValidateJWT(tokenString)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            ctx.Abort()
            return
        }

        ctx.Set("email", claims.Email)
        ctx.Next()
	}
}