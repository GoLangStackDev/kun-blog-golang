package cmd

import (
	"encoding/base64"
	"encoding/pem"
	"github.com/spf13/cobra"
	"io/ioutil"
	"kun-blog-golang/pkg/kctl/config"
	"log"
)

var ca string
var clientKey string
var clientCert string

func init() {
	configCMD.Flags().StringVar(&clientCert, "cert", "", "设置客户端证书")
	configCMD.Flags().StringVar(&clientKey, "key", "", "设置客户端证书KEY")
	configCMD.Flags().StringVar(&ca, "ca", "", "设置CA证书")
}

func loadCa(f string) string {
	r, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	block, _ := pem.Decode(r) //ca特有的解码
	return base64.StdEncoding.EncodeToString(block.Bytes)
}

func loadCertOrKey(f string) string {
	r, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	return base64.StdEncoding.EncodeToString(r)
}

var configCMD = &cobra.Command{
	Use:          "config",
	Example:      "kub config",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		cfg := config.LoadConfigFile()
		isUpdate := false
		if clientCert != "" {
			isUpdate = true
			cfg.ClientCertData = loadCertOrKey(clientCert)
		}
		if clientKey != "" {
			isUpdate = true
			cfg.ClientKeyData = loadCertOrKey(clientKey)
		}
		if ca != "" {
			isUpdate = true
			cfg.CaData = loadCa(ca)
		}
		if isUpdate {
			config.SaveConfig(cfg)
		} else {
			log.Println("请设置 --client-cert --client-key --ca 这三个参数")
		}
		return nil
	},
}
