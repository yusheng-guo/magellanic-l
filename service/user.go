package service

import (
	"github.com/yushengguo557/magellanic-l/common/request"
)

type UserServiceT struct {
}

var UserService = new(UserServiceT)

// Login 登录
func (s *UserServiceT) Login(params request.Login) (string, error) {
	//db := global.App.MongoDB
	return "", nil
}

func (s *UserServiceT) Register(params request.Register) (string, error) {
	//db := global.App.MongoDB
	return "", nil
}
