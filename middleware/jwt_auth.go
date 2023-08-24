package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/service"
	"net/http"
)

// JwtAuth JWT
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "request does not contain an access token"})
			c.Abort()
			return
		}
		err := service.JwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
