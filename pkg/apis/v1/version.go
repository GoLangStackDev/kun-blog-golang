package v1

import "encoding/json"

type Resolve struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// MustDataByte 返回 data 的 []byte 类型
func (this *Resolve) MustDataByte() []byte {
	rst, err := json.Marshal(this.Data)
	if err != nil {
		panic(err)
	}
	return rst
}

func NewResolve(data interface{}) *Resolve {
	return &Resolve{Msg: "操作成功", Code: 200, Data: data}
}

type Version struct {
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
}
