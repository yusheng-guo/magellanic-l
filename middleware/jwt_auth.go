package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/service"
	"net/http"
	"strings"
)

const AuthSchemeBearer = "bearer" // 验证方案

// JwtAuth JWT
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization") // Authorization: <auth-scheme> <authorization-parameters>
		params := strings.Split(strings.TrimSpace(auth), " ")
		if len(params) > 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "incorrect token format"})
			c.Abort()
			return
		}
		// 身份验证方案 auth-scheme 不区分大小写
		if strings.ToLower(params[0]) != AuthSchemeBearer {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "auth scheme incorrect"})
			c.Abort()
			return
		}
		err := service.JwtService.ValidateToken(params[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
