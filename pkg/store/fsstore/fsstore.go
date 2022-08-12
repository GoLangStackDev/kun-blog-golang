package fsstore

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"kun-blog-golang/pkg/store"
	"kun-blog-golang/pkg/store/models"
	"kun-blog-golang/pkg/utils"
	"log"
	"os"
	"path/filepath"
)

type PostsMap map[string]*models.Post

func (this PostsMap) All() (rst []*models.Post) {
	rst = make([]*models.Post, len(this))
	i := 0
	for _, post := range this {
		rst[i] = post
		i++
	}
	return
}

type FSStore struct {
	root     string
	catchMap PostsMap
}

func NewFSStore(root string) *FSStore {
	rs := &FSStore{
		root:     root,
		catchMap: make(PostsMap),
	}
	if err := rs.catchFiles(); err != nil {
		log.Println(err)
	}
	return rs
}

func (this *FSStore) catchFiles() (err error) {
	return filepath.Walk(this.root, func(path string, info fs.FileInfo, err error) error {
		var f *os.File
		var postM *models.Post
		// 处理文件异常以及是文件夹的情况
		if fileStat, err := os.Stat(path); err != nil {
			return err
		} else if fileStat.IsDir() {
			return nil
		}
		// 打开一个文件句柄
		if f, err = os.Open(path); err != nil {
			return err
		}
		defer f.Close() //记得关闭
		// 解析
		if postM, err = utils.ParseFileToPost(f); err != nil {
			return err
		}
		// 保存到缓存
		this.catchMap[postM.Slug] = postM
		return nil
	})
}

// 声明实现接口
var _ store.StoreInterface = &FSStore{}

func (this *FSStore) List() (posts []*models.Post, err error) {
	return this.catchMap.All(), nil
}

func (F FSStore) DeleteBySlug(slug string) error {
	//TODO implement me
	panic("implement me")
}

func (this *FSStore) GetBySlug(slug string) (*models.Post, error) {
	return this.catchMap[slug], nil
}

func (this *FSStore) Save(post *models.Post) error {
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
	filePath := fmt.Sprintf("%s/%s", this.root, post.FileName)

	// 加入缓存
	this.catchMap[post.Slug] = post

	return os.WriteFile(filePath, mdBuffer.Bytes(), 0600)
}
