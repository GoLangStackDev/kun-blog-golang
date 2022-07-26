package clientset

import (
	"kun-blog-golang/pkg/kctl/clientset/rest"
	"testing"
	"time"
)

var cfg = &rest.Config{
	Host:    "http://www.baidu.com",
	TimeOut: time.Second * 10,
}

func TestV1(t *testing.T) {
	client := NewClientSetForConfig(cfg)
	client.V1().Version().Get() //链式调用
}
