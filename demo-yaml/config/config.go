package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Configs 整合所有的 struct
type Configs struct {
	Database Database `map structure:"database"`
	Server   Server   `map structure:"server"`
}

// Database 数据库相关
type Database struct {
	Host     string `map structure:"host" json:"host" yaml:"host"`
	Port     int    `map structure:"port" json:"port" yaml:"port"`
	Name     string `map structure:"name" json:"name" yaml:"name"`
	User     string `map structure:"user" json:"user" yaml:"user"`
	Password string `map structure:"password" json:"password" yaml:"password"`
}

// Server 服务相关
type Server struct {
	Port int `map structure:"port" yaml:"port" json:"port" `
}

// GlobalConfig 定义全局变量
var GlobalConfig Configs

// Load 读取config 文件下的配置 相比spring boot 的 application.yaml 可以直接读取 还是差一丢丢体验
func Load() {
	configFile, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}

	err = yaml.Unmarshal(configFile, &GlobalConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
}
