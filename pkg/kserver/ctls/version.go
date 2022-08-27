package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	v1 "kun-blog-golang/pkg/apis/v1"
)

type VersionCtl struct{}

func NewVersionCtl() *VersionCtl {
	return &VersionCtl{}
}

func (this *VersionCtl) Version(c *gin.Context) {
	rst := &v1.Resolve{
		Code: 200,
		Msg:  "操作成功",
		Data: &v1.Version{
			Version:   "v0.1",
			GoVersion: "go1.17",
		},
	}
	c.JSON(200, rst)
}

// 声明需要实现 IClass 这个接口
var _ server.IClass = &VersionCtl{}

// Build 实现 IClass 接口必要方法
func (this *VersionCtl) Build(core *server.KCore) {
	core.Handle("GET", "/version", this.Version) //支持
}
