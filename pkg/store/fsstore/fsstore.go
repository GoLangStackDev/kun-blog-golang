package fsstore

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"kun-blog-golang/pkg/store"
	"kun-blog-golang/pkg/store/models"
	"os"
)

type FSStore struct {
	root string
}

func NewFSStore(root string) *FSStore {
	rs := &FSStore{root: root}
	return rs
}

// 声明实现接口
var _ store.StoreInterface = &FSStore{}

func (F FSStore) List() []*models.Post {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) DeleteBySlug(slug string) error {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) GetBySlug(slug string) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) Save(post *models.Post) error {
	var err error

	// 解析成 yaml byte
	b, err := yaml.Marshal(post.PostConfig)
	if err != nil {
		return err
	}

	// 缓冲的buffer
	var mdBuffer bytes.Buffer
	// 写入特殊标识
	mdBuffer.WriteString("---\n")
	// 写入配置字符
	mdBuffer.WriteString(string(b))
	// 写入特殊标识
	mdBuffer.WriteString("---\n")
	// 写入markdown主体内容
	mdBuffer.WriteString(post.Md)

	// 写入文件
	return os.WriteFile(post.FilePath, mdBuffer.Bytes(), 0600)
}
