package clientset

import (
	"encoding/base64"
	"kun-blog-golang/pkg/kctl/clientset/rest"
	v1 "kun-blog-golang/pkg/kctl/clientset/typed/v1"
	"kun-blog-golang/pkg/kctl/config"
	"log"
	"time"
)

var DefClientSet *ClientSet

func init() {
	sysCfg := config.LoadConfigFile()
	cfg := &rest.Config{
		Host:    sysCfg.ServerHost,
		TimeOut: time.Duration(sysCfg.TimeOutSecond) * time.Second,
		// ssl
		ClientCertData: Base64Decode(sysCfg.ClientCertData),
		ClientKeyData:  Base64Decode(sysCfg.ClientKeyData),
		CaData:         Base64Decode(sysCfg.CaData),
	}
	DefClientSet = NewClientSetForConfig(cfg)
}

// Base64Decode 解码证书
func Base64Decode(str string) []byte {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}

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
