package main

import (
	"kun-blog-golang/core/server"
	"kun-blog-golang/pkg/kserver/ctls"
)

func main() {
	r := server.New()
	r.Mount(
		ctls.NewVersionCtl())
	r.Run(":8080")
}
