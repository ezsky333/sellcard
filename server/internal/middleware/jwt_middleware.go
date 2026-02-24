package middleware

import (
    "net/http"
    "strings"
    "sellcard/server/internal/config"
    "sellcard/server/internal/utils"

    "github.com/gin-gonic/gin"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
            return
        }
        parts := strings.SplitN(auth, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
            return
        }
        token := parts[1]
        claims, err := utils.ParseToken(cfg.JWT.Secret, token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        c.Set("claims", claims)
        c.Next()
    }
}
