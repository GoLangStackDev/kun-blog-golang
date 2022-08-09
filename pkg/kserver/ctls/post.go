package ctls

import (
	"github.com/gin-gonic/gin"
	"kun-blog-golang/core/server"
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
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
		"data": "",
	})
}

func (this *PostCtl) Build(core *server.KCore) {
	core.Handle("POST", "/post/apply", this.Apply)
}
