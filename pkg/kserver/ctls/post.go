package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
	v1 "kun-blog-golang/pkg/apis/v1"
	"kun-blog-golang/pkg/store/fsstore"
	. "kun-blog-golang/pkg/utils"
	"log"
)

type PostCtl struct {
	fs          *fsstore.FSStore
	articlePath string
}

func NewPostCtl() *PostCtl {
	articlePath := "./public/contents/articles"
	return &PostCtl{
		fs:          fsstore.NewFSStore(articlePath),
		articlePath: articlePath,
	}
}
func (this *PostCtl) Apply(c *gin.Context) {
	// 从请求里面取表单
	frm, err := c.MultipartForm()
	CheckError(err)
	// 打印下取到的表单信息
	log.Println("apply", frm)
	for inputName, fileList := range frm.File { //支持多文件上传
		log.Println("处理文件:", inputName)
		f, err := fileList[0].Open()
		CheckError(err)
		defer f.Close() //记得关闭
		postM, err := ParseFileToPost(f)
		CheckError(err)
		postM.FileName = fileList[0].Filename //赋值文件名
		CheckError(this.fs.Save(postM))
	}
	c.JSON(200, v1.Resolve{
		Code: 200,
		Msg:  "操作成功",
	})
}

func (this *PostCtl) Get(c *gin.Context) {
	slug := c.Param("slug")
	m, err := this.fs.GetBySlug(slug)
	CheckError(err)
	if m == nil {
		c.JSON(200, v1.Resolve{
			Code: 404,
			Msg:  "没找到文章",
		})
	} else {
		c.JSON(200, v1.Resolve{
			Code: 200,
			Msg:  "操作成功",
			Data: m,
		})
	}
}

func (this *PostCtl) Delete(c *gin.Context) {
	slug := c.Param("slug")
	CheckError(this.fs.DeleteBySlug(slug))
	c.JSON(200, v1.Resolve{
		Code: 200,
		Msg:  "操作成功",
	})
}

func (this *PostCtl) List(c *gin.Context) {
	list, err := this.fs.List()
	CheckError(err)
	c.JSON(200, v1.Resolve{
		Code: 200,
		Msg:  "操作成功",
		Data: list,
	})
}

func (this *PostCtl) Build(core *server.KCore) {
	core.Handle("POST", "/posts/apply", this.Apply)
	core.Handle("GET", "/posts/:slug", this.Get)
	core.Handle("DELETE", "/posts/:slug", this.Delete)
	core.Handle("GET", "/posts", this.List)
}
