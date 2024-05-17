package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config 定义一个结构体
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/aaa/v")
	var config Config
	err := viper.Unmarshal(&config)
	fmt.Println("\n111111:\n", err)
	if err != nil {
		fmt.Println("\n读取日志失败:\n", err)
		return
	}
	fmt.Println("\n输出User:\n", config)
}
