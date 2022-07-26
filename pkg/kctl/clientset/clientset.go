package clientset

import (
	"kun-blog-golang/pkg/kctl/clientset/rest"
	v1 "kun-blog-golang/pkg/kctl/clientset/typed/v1"
)

type ClientSet struct {
	*rest.RESTClient
}

func NewClientSetForConfig(c *rest.Config) *ClientSet {
	return &ClientSet{
		RESTClient: rest.NewRESTClient(c),
	}
}

func (this *ClientSet) V1() v1.V1Interface {
	return v1.New(this.RESTClient)
}
