package rest

import (
	"fmt"
	"net/http"
	"net/url"
)

type Request struct {
	c       *RESTClient
	path    string
	params  url.Values
	headers http.Header
	method  string
}

func NewRequest(c *RESTClient) *Request {
	return &Request{c: c, path: "/"}
}

func (this *Request) Method(method string) *Request {
	this.method = method
	return this
}

func (this *Request) Path(p string) *Request {
	this.path = p
	return this
}

func (this *Request) Do() *Result {
	rest := &Result{}
	rsp, err := this.c.R().Execute(this.method, this.path)
	if err != nil {
		rest.err = err
	} else if rsp.IsError() {
		if rsp.Error() == nil { //兼容 http code 不是200
			//rest.err = errors.New("HTTP Status Code Is Not 200")
			rest.rsp = rsp
		} else {
			rest.err = fmt.Errorf("%v", rsp.Error())
		}
	} else {
		rest.rsp = rsp
	}
	return rest
}
