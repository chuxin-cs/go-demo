package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Hello `yaml:"hello"`
	Redis `yaml:"redis"`
}

type Hello struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

func main() {

	configFile, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}

	fmt.Println("yaml 文件的内容:\n", string(configFile))

	config := Config{}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}

	host := config.Host

	fmt.Printf("\n host内容为:\n", host)
	fmt.Printf("\nconfig内容为:\n", config)
	fmt.Printf("\nHost==============>: %s\n", config.Redis.Host)

	//mp := make(map[string]any, 2)
	//err = yaml.Unmarshal(configFile, mp)
	//if err != nil {
	//	fmt.Println("解析 yaml 文件失败：", err)
	//	return
	//}
	//fmt.Println("内容为:", mp)
}
