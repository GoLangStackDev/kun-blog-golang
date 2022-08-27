package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	// 服务端地址
	ServerHost string `yaml:"server_host"`
	// 连接超时时间
	TimeOutSecond int `yaml:"time_out_second"`

	// ssl证书
	ClientCertData string `yaml:"client_cert_data"`
	ClientKeyData  string `yaml:"client_key_data"`
	CaData         string `yaml:"ca_data"`
}

func getConfigFilePath() string {
	home, err := os.UserHomeDir() //获取用的家目录地址
	if err != nil {
		log.Fatalln(err)
	}
	configFile := home + "/.kblog/config" //拼接出配置文件地址
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		log.Fatalln(configFile, "配置文件没找到")
	}
	return configFile
}

func LoadConfigFile() *Config {
	configFilePath := getConfigFilePath()
	cfg := &Config{}
	err := yaml.Unmarshal(MustLoadFile(configFilePath), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}

func MustLoadFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

func SaveConfig(cfg *Config) {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	path := getConfigFilePath()
	err = os.WriteFile(path, b, 0600)
	if err != nil {
		log.Fatalln(err)
	}
}
