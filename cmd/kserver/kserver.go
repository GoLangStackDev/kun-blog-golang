package main

import (
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/kserver/ctls"
	"kun-blog-golang/pkg/utils"
)

func main() {
	r := server.New()
	// 设置静态文件资源地址
	r.Static("/statics", "./public/statics")

	// 注入模板方法
	r.SetFuncMap(utils.HTMLFuncMap)

	// 设置加载的html模板地址
	r.LoadHTMLGlob("./public/tpl/*")
	r.Mount("", ctls.NewPageCtl()) //添加页面控制器
	r.Mount(
		"/v1",
		ctls.NewVersionCtl(),
		ctls.NewPostCtl())

	r.Start()
}
