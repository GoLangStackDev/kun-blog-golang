package fsstore

import (
	"kun-blog-golang/pkg/store/models"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	fs := NewFSStore("")
	err := fs.Save(&models.Post{
		PostConfig: &models.PostConfig{
			Slug:        "ABC123",
			Title:       "第一篇文章",
			CreatedTime: time.Now(),
			Author:      "小锟哥哥",
			Describe:    "文章描述",
		},
		Md:       "# 标题",
		FilePath: "./hello.md",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("保存成功")
}
