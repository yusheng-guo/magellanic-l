package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yushengguo557/magellanic-l/internal/service"
)

func WebSocket(c *gin.Context) {
	//uid, exists := c.Get("uid")
	//if !exists {
	//	response.Failed(c, errors.New("unable to get user id from context"))
	//}
	uid := uuid.NewString()
	service.WebSocketHandel(uid, c.Writer, c.Request)
}
