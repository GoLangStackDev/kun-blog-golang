package rest

import (
	"github.com/go-resty/resty/v2"
	"net/http"
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

func (this *RESTClient) Get() *Request {
	return NewRequest(this).Method(http.MethodGet)
}
func (this *RESTClient) Post() *Request {
	return NewRequest(this).Method(http.MethodPost)
}
func (this *RESTClient) Delete() *Request {
	return NewRequest(this).Method(http.MethodDelete)
}
