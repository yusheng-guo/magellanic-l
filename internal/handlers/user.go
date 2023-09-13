package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/api"
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/internal/service"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	var err error
	var form api.LoginReq
	var token string

	if err = c.ShouldBindQuery(&form); err != nil {
		global.App.Log.Error("bind params", zap.Any("err", err))
		api.RequestErrorHandler(c, errors.New("params error"))
		//response.Failed(c, response.ParamsError, "Please enter email and password.")
		return
	}

	token, err = service.UserService.Login(form.Email, form.Password)
	if err != nil {
		global.App.Log.Error("user login", zap.Any("err", err))
		api.InternalErrorHandler(c)
		//response.Failed(c, response.InternalError, "Please try again later.")
		return
	}
	//response.Succeed(c, token)
	api.SuccessHandler(c, token)
}

func Register(c *gin.Context) {}

func ObtainInfo(c *gin.Context) {}

func EditInfo(c *gin.Context) {}
