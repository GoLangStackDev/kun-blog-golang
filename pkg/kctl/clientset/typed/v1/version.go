package v1

import (
	v1 "kun-blog-golang/pkg/apis/v1"
	"kun-blog-golang/pkg/kctl/clientset/rest"
)

// VersionInterface 定义 /version 下面有那些方法
type VersionInterface interface {
	Get() (rVersion *v1.Resolve, err error)
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
func (this *version) Get() (rVersion *v1.Resolve, err error) {
	rVersion = &v1.Resolve{}
	err = this.client.Get().Path("/v1/version").Do().Into(rVersion)
	return
}
