package clientset

import (
	"fmt"
	"kun-blog-golang/pkg/kctl/clientset/rest"
	"testing"
	"time"
)

var cfg = &rest.Config{
	Host:    "http://127.0.0.1:8080",
	TimeOut: time.Second * 10,
}

func TestV1(t *testing.T) {
	client := NewClientSetForConfig(cfg)
	ver, err := client.V1().Version().Get() //链式调用
	fmt.Println(ver, err)
}
