package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Groups []string `yaml:"groups"`
}

func (c *Config) String() string {
	sb := strings.Builder{}
	sb.WriteString("\n[\n")
	for idx, group := range c.Groups {
		sb.WriteString(fmt.Sprintf("  %d -> %s\n", idx+1, group))
	}
	sb.WriteString("]\n")
	return sb.String()
}

var C *Config

// Init 读取配置文件
func Init() error {
	cfgBytes, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}
	C = new(Config)
	if err := yaml.Unmarshal(cfgBytes, C); err != nil {
		return err
	}
	log.Println("配置加载成功: ", C)
	return nil
}
