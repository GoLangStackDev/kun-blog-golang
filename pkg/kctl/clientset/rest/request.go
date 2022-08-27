package rest

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
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

func (this *Request) ApplyFileByte(url string, fileName string, fileByte []byte, isForce bool) *Result {
	req := this.c.R().
		SetHeader(
			"Content-Length",
			fmt.Sprintf("%v", len(fileByte))).
		SetFileReader(fileName, fileName, bytes.NewReader(fileByte))
	if isForce {
		req.SetHeader("Force", fmt.Sprintf("%v", isForce))
	}
	return toResult(req.Post(url))
}

func (this *Request) Do() *Result {
	return toResult(this.c.R().Execute(this.method, this.path))
}

func toResult(rsp *resty.Response, err error) *Result {
	rest := &Result{}
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
