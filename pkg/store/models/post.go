package models

import "time"

type PostConfig struct {
	Slug        string    `yaml:"slug"`
	Title       string    `yaml:"title"`
	CreatedTime time.Time `yaml:"created_time"`
	Author      string    `yaml:"author"`
	Describe    string    `yaml:"describe"`
}

type Post struct {
	*PostConfig
	FileName  string
	Html      string
	Md        string
	LocalPath string //本地地址
}
