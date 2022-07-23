package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type KCore struct {
	*gin.Engine
	g *gin.RouterGroup
}

func New() *KCore {
	return &KCore{
		Engine: gin.New(), //这里一定要初始化
	}
}

// Mount 挂载我们的控制器
func (this *KCore) Mount(group string, classes ...IClass) *KCore {
	this.g = this.Group(group)
	for _, class := range classes {
		// 把自己传进去
		class.Build(this)
	}
	return this
}

// Handle 接管路由挂载
func (this *KCore) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *KCore {
	// 一定基于我们自己的路由组进行挂载
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

func (this *KCore) Start() {
	var port int32 = 8080
	this.Run(fmt.Sprintf(":%d", port))
}
