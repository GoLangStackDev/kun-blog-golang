package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/store/fsstore"
	"log"
	"net/http"
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
	list, err := this.fs.List()
	if err != nil {
		log.Println(err)
	}
	c.HTML(200, "index.html", gin.H{
		"Posts": list,
		"Title": "首页",
	})
}

// DetailPage 详情页
func (this *PageCtl) DetailPage(c *gin.Context) {
	slug := c.Param("slug")
	post, err := this.fs.GetBySlug(slug)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"Title": "详情",
		"Post":  post,
	})
}

func (this *PageCtl) Build(core *server.KCore) {
	core.Handle("GET", "/", this.HomePage)
	core.Handle("GET", "/posts/:slug", this.DetailPage)
}
