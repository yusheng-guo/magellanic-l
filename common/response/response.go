package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusCode uint8

const (
	Success StatusCode = iota
	ParamsError
	InternalError
)

type Response struct {
	Code StatusCode `json:"code"`
	Msg  string     `json:"msg"`
	Data any        `json:"data"`
}

// Failed 失败响应
func Failed(c *gin.Context, code StatusCode, msg string) {
	c.JSON(http.StatusOK,
		Response{
			Code: code,
			Msg:  msg,
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
