package router

import (
	"github.com/yushengguo557/magellanic-l/global"
)

func DefineRouter() {
	group := NewGroup(global.App.Engine.RouterGroup)
	group.AddPing()
	api := group.SetGroup("/api")
	v1 := api.SetGroup("/v1")
	{
		v1.AddPing()

		v1.AddAuth()

		v1.AddWebSocket()
	}

}
