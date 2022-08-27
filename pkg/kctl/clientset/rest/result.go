package rest

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Result struct {
	rsp *resty.Response
	err error
}

func (this *Result) Into(v interface{}) (err error) {
	if this.err != nil {
		return this.err
	}
	//log.Println(string(this.rsp.Body()))
	err = json.Unmarshal(this.rsp.Body(), v)
	if err != nil {
		return
	}
	return
}
