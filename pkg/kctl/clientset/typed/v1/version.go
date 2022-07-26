package v1

import (
	"fmt"
	"kun-blog-golang/pkg/kctl/clientset/rest"
)

// VersionInterface 定义 /version 下面有那些方法
type VersionInterface interface {
	Get()
}

type version struct {
	client *rest.RESTClient
}

func newVersion(client *rest.RESTClient) *version {
	return &version{client: client}
}

// 声明实现接口
var _ VersionInterface = &version{}

// Get 实现 VersionInterface 接口
func (v version) Get() {
	fmt.Println("version get")
}
