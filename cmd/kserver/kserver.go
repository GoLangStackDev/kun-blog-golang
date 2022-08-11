package main

import (
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/kserver/ctls"
)

func main() {
	r := server.New()
	// 设置加载的html模板地址
	r.LoadHTMLGlob("./public/tpl/*")
	r.Mount("", ctls.NewPageCtl()) //添加页面控制器
	r.Mount(
		"/v1",
		ctls.NewVersionCtl(),
		ctls.NewPostCtl())

	r.Start()
}
