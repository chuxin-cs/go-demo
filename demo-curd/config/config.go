package config

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File

	HTTPPort int
)

func init() {
	// 定义变量
	var err error

	// 读取配置
	Cfg, err = ini.Load("conf/app.ini")

	// 读取的时候如果不是 nil 说明有错误
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadServer()
}

func LoadServer() {

	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(9000)
}
