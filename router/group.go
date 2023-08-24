package router

import "github.com/gin-gonic/gin"

// Group 自定义路由组 封装了 Gin 框架的 RouterGroup (相当于继承)
type Group struct {
	gin.RouterGroup
}

// NewGroup 实例化一个组
func NewGroup(g gin.RouterGroup) *Group {
	return &Group{
		g,
	}
}

// SetGroup 设置组
func (g *Group) SetGroup(relativePath string, handlers ...gin.HandlerFunc) *Group {
	return &Group{
		*g.Group(relativePath, handlers...),
	}
}
