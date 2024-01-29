package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Database struct {
	Driver       string `yaml:"driver"`        // 驱动名称
	User         string `yaml:"user"`          // 用户名
	Password     string `yaml:"password"`      // 密码
	Host         string `yaml:"host"`          // 主机地址
	Port         int    `yaml:"port"`          // 端口
	Name         string `yaml:"name"`          // 库名
	Prefix       string `yaml:"prefix"`        // 表名前缀
	Dns          string `yaml:"dns"`           // 连接dns
	MaxIdleConns int    `yaml:"maxIdle_conns"` // 最大空闲数
	MaxOpenConns int    `yaml:"maxOpen_conns"` // 最大连接数
}

type Debug struct {
	Enable bool `yaml:"enable"`
}

type Config struct {
	RouterPrefix string               `yaml:"router_prefix"` // 路由前缀
	TemplateName string               `yaml:"template_name"` // 模板名称
	ServerListen string               `yaml:"server_listen"` // http服务配置
	Databases    map[string]*Database `yaml:"databases"`     // db配置
	Debug        *Debug               `yaml:"debug"`         // 调试模式
}

var config = &Config{}

// LoadFromYaml 加载yaml配置
func LoadFromYaml(path string) (*Config, error) {
	jsonByte, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(jsonByte, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GetDebugConf 获取调试配置
func GetDebugConf() *Debug {
	return config.Debug
}

// GetTemplateName 获取调试配置
func GetTemplateName() string {
	return config.TemplateName
}
