package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	cfg := LoadConfigFile()
	fmt.Println(cfg)
}
