package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/common/request"
	"github.com/yushengguo557/magellanic-l/common/response"
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/service"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	var err error
	var form request.Login
	var token string

	if err = c.ShouldBindQuery(&form); err != nil {
		global.App.Log.Error("bind params", zap.Any("err", err))
		response.Failed(c, response.ParamsError, "Please enter email and password.")
	}

	token, err = service.UserService.Login(form)
	if err != nil {
		global.App.Log.Error("user login", zap.Any("err", err))
		response.Failed(c, response.InternalError, "Please try again later.")
	}
	response.Succeed(c, token)
}

func Register(c *gin.Context) {}

func ObtainInfo(c *gin.Context) {}

func EditInfo(c *gin.Context) {}
