package rest

import (
	"github.com/go-resty/resty/v2"
)

type RESTClient struct {
	*resty.Client
}

func NewRESTClient(cfg *Config) *RESTClient {
	rc := resty.New()
	rc.SetBaseURL(cfg.Host)    //请求的基础URL
	rc.SetTimeout(cfg.TimeOut) //请求超时时间
	return &RESTClient{Client: rc}
}
