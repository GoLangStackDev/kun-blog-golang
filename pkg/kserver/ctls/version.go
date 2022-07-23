package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
)

type VersionCtl struct{}

func NewVersionCtl() *VersionCtl {
	return &VersionCtl{}
}

func (this *VersionCtl) Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"Version":   "v0.1",
		"GoVersion": "go1.7",
	})
}

func (this *VersionCtl) Build(core *server.KCore) {
	core.Handle("GET", "/version", this.Version) //支持
}
