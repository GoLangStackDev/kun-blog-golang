package main

import (
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/kserver/ctls"
)

func main() {
	r := server.New()
	r.Mount(
		"/v1",
		ctls.NewVersionCtl(),
		ctls.NewPostCtl()) //新增

	r.Start()
}
