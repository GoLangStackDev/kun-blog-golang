package v1

import (
	"errors"
	"fmt"
	"kun-blog-golang/pkg/apis/code"
	v1 "kun-blog-golang/pkg/apis/v1"
	"kun-blog-golang/pkg/kctl/clientset/rest"
)

type Post struct {
	client *rest.RESTClient
}

func NewPost(client *rest.RESTClient) *Post {
	return &Post{client: client}
}

// PostsInterface Post对应的API
type PostsInterface interface {
	Apply(fileName string, fileByte []byte, isForce bool) (rst *v1.Resolve, err error)
	List() (posts []*v1.Post, err error)
	GetBySlug(slug string) (post *v1.Post, err error)
	DeleteBySlug(slug string) (err error)
}

// 实现接口
var _ PostsInterface = &Post{}

func (this *Post) List() (posts []*v1.Post, err error) {
	rst := &v1.PostListResolve{}
	err = this.client.Get().Path("/v1/posts/list").Do().Into(rst)
	if err != nil {
		return nil, err
	}
	if rst.Code != code.Success {
		return nil, errors.New(rst.Msg)
	}
	return rst.Data, nil
}

func (this *Post) GetBySlug(slug string) (post *v1.Post, err error) {
	rst := &v1.PostResolve{}
	err = this.client.Get().Path(fmt.Sprintf("/v1/posts/%s", slug)).Do().Into(rst)
	if err != nil {
		return nil, err
	}
	if rst.Code != code.Success {
		return nil, errors.New(rst.Msg)
	}
	return rst.Data, nil
}

func (this *Post) DeleteBySlug(slug string) (err error) {
	rst := &v1.Resolve{}
	err = this.client.Delete().Path(fmt.Sprintf("/v1/posts/%s", slug)).Do().Into(rst)
	if err != nil {
		return err
	}
	if rst.Code != code.Success {
		return errors.New(rst.Msg)
	}
	return nil
}

func (this *Post) Apply(fileName string, fileByte []byte, isForce bool) (rst *v1.Resolve, err error) {
	rst = &v1.Resolve{}
	err = this.client.Post().
		ApplyFileByte("/v1/posts/apply", fileName, fileByte, isForce).
		Into(rst)
	return
}
