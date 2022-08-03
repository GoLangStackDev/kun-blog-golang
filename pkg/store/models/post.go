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
	Html     string
	Md       string
	FilePath string `yaml:"-"`
}
