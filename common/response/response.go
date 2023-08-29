package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusCode uint8

const (
	Success StatusCode = iota
)

type Response struct {
	Code StatusCode `json:"code"`
	Msg  string     `json:"msg"`
	Data any        `json:"data"`
}

// Failed 失败响应
func Failed(c *gin.Context, err error) {
	c.JSON(http.StatusOK,
		Response{
			Code: Success,
			Msg:  err.Error(),
			Data: nil,
		},
	)
}

// Succeed 成功响应
func Succeed(c *gin.Context, data any) {
	c.JSON(http.StatusOK,
		Response{
			Code: Success,
			Msg:  "",
			Data: data,
		},
	)
}
