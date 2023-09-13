package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

var (
	RequestErrorHandler = func(c *gin.Context, err error) {
		writeResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	InternalErrorHandler = func(c *gin.Context) {
		writeResponse(c, http.StatusInternalServerError, "An Unexpected Error Occurred", nil)
	}
	SuccessHandler = func(c *gin.Context, data any) {
		writeResponse(c, http.StatusOK, "", data)
	}
)

func writeResponse(c *gin.Context, code int, message string, data any) {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}
