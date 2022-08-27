package utils

import (
	"io/ioutil"
	"os"
)

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func LoadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
func MustLoadFile(path string) []byte {
	b, err := LoadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
func FileInfo(path string) (os.FileInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f.Stat()
}
