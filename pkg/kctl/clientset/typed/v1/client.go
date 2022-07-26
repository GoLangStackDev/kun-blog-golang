package v1

import "kun-blog-golang/pkg/kctl/clientset/rest"

type V1Client struct {
	client *rest.RESTClient
}

func New(client *rest.RESTClient) *V1Client {
	return &V1Client{client: client}
}

// 声明实现 V1Interface 接口
var _ V1Interface = &V1Client{}

// 声明实现 version 接口
var _ VersionGetter = &V1Client{}

// Version 实现 VersionGetter 接口的方法
func (this *V1Client) Version() VersionInterface {
	return newVersion(this.client)
}
