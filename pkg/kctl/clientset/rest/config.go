package rest

import "time"

type Config struct {
	Host    string
	TimeOut time.Duration

	// ssl证书
	ClientCertData []byte
	ClientKeyData  []byte
	CaData         []byte
}
