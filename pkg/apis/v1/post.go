package v1

import "time"

type Post struct {
	Title       string
	Slug        string
	Md          string
	Html        string
	CreatedTime time.Time `yaml:"created_time"`
	Author      string    `yaml:"author"`
}

type PostListResolve struct {
	Msg  string  `json:"msg"`
	Code int     `json:"code"`
	Data []*Post `json:"data"`
}

type PostResolve struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data *Post  `json:"data"`
}
