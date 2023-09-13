package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/api"
	"github.com/yushengguo557/magellanic-l/internal/service"
	"strings"
)

const AuthSchemeBearer = "bearer" // 验证方案

var (
	TokenFormatError = errors.New("incorrect token format")
	TokenSchemeError = errors.New("auth scheme incorrect")
)

// UnAuthorizedError = errors.New("invalid token")

// JwtAuth JWT
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization") // Authorization: <auth-scheme> <authorization-parameters>
		params := strings.Split(strings.TrimSpace(auth), " ")
		if len(params) > 2 {
			//c.JSON(http.StatusUnauthorized, gin.H{"err": "incorrect token format"})
			api.RequestErrorHandler(c, TokenFormatError)
			c.Abort()
			return
		}
		// 身份验证方案 auth-scheme 不区分大小写
		if strings.ToLower(params[0]) != AuthSchemeBearer {
			//c.JSON(http.StatusUnauthorized, gin.H{"err": "auth scheme incorrect"})
			api.RequestErrorHandler(c, TokenSchemeError)
			c.Abort()
			return
		}
		err := service.JwtService.ValidateToken(params[1])
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			api.RequestErrorHandler(c, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
