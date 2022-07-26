package rest

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

var cfg = &Config{
	Host:    "http://www.baidu.com",
	TimeOut: time.Second * 10,
}

func TestNewRESTClient(t *testing.T) {
	req := NewRESTClient(cfg)
	rsp, err := req.R().Execute(http.MethodGet, "/")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(string(rsp.Body()))
	}
}
