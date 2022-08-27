package rest

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type RESTClient struct {
	*resty.Client
}

func NewRESTClient(cfg *Config) *RESTClient {
	rc := resty.New()
	rc.SetBaseURL(cfg.Host)    //请求的基础URL
	rc.SetTimeout(cfg.TimeOut) //请求超时时间

	// 设置客户端证书
	clientPEM, err := tls.X509KeyPair(cfg.ClientCertData, cfg.ClientKeyData)
	if err != nil {
		log.Fatalln(err)
	}
	// 设置 tls 证书
	certPool := x509.NewCertPool()
	caPEM, _ := x509.ParseCertificate(cfg.CaData)
	certPool.AddCert(caPEM)
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientPEM},
	}
	rc.SetTLSClientConfig(tlsConfig)

	return &RESTClient{Client: rc}
}

func (this *RESTClient) Get() *Request {
	return NewRequest(this).Method(http.MethodGet)
}
func (this *RESTClient) Post() *Request {
	return NewRequest(this).Method(http.MethodPost)
}
func (this *RESTClient) Delete() *Request {
	return NewRequest(this).Method(http.MethodDelete)
}
