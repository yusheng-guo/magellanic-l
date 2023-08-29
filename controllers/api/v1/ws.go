package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/common/response"
	"github.com/yushengguo557/magellanic-l/service/ws"
)

func WebSocket(c *gin.Context) {
	//handler := websocket.Handler(ws.Echo)
	uid, exists := c.Get("uid")
	if !exists {
		response.Failed(c, errors.New("unable to get user id from context"))
	}
	ws.WebSocketHandel(uid.(string), c.Writer, c.Request)

}
