package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	v1 "kun-blog-golang/pkg/apis/v1"
)

type Result struct {
	rsp *resty.Response
	err error
}

func (this *Result) Into(v interface{}) (err error) {
	if this.err != nil {
		return this.err
	}
	rst := &v1.Resolve{}
	err = json.Unmarshal(this.rsp.Body(), rst)
	if err != nil {
		return
	}
	if rst.Code != 200 { // 通用代码处理，非 200 统一报错
		err = fmt.Errorf("code: %d,msg: %s", rst.Code, rst.Msg)
		return
	}
	err = json.Unmarshal(rst.MustDataByte(), v)
	if err != nil {
		return
	}
	return
}
