package server

import "github.com/gin-gonic/gin"

type KCore struct {
	*gin.Engine
}

func New() *KCore {
	return &KCore{
		Engine: gin.New(), //这里一定要初始化
	}
}

// Mount 挂载我们的控制器
func (this *KCore) Mount(classes ...IClass) *KCore {
	for _, class := range classes {
		// 把自己传进去
		class.Build(this)
	}
	return this
}
