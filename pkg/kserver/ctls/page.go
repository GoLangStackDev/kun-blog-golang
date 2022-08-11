package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/store/fsstore"
	"log"
)

type PageCtl struct {
	fs          *fsstore.FSStore
	articlePath string
}

func NewPageCtl() *PageCtl {
	articlePath := "./public/contents/articles"
	return &PageCtl{
		fs:          fsstore.NewFSStore(articlePath),
		articlePath: articlePath,
	}
}

func (this *PageCtl) HomePage(c *gin.Context) {
	//list := make([]*models.Post, 0)
	list, err := this.fs.List()
	if err != nil {
		log.Println(err)
	}
	c.HTML(200, "index.html", gin.H{
		"Posts": list,
		"Title": "首页",
	})
}

func (this *PageCtl) Build(core *server.KCore) {
	core.Handle("GET", "/", this.HomePage)
}
