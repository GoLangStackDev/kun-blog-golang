package utils

import (
	"html/template"
	"time"
)

// HTMLFuncMap html模板里面的自定义方法
var HTMLFuncMap map[string]interface{}

func init() {
	HTMLFuncMap = make(map[string]interface{})
	HTMLFuncMap["Html"] = Html
	HTMLFuncMap["TimeFormat"] = TimeFormat
}

// Html 转换成html显示
func Html(input string) interface{} {
	return template.HTML(input)
}

// TimeFormat 格式化时间输出
func TimeFormat(input time.Time, format string) string {
	if format == "" {
		format = "2006-01-02"
	}
	return input.Format(format)
}
