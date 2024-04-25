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
	var err error

	Cfg, err = ini.Load("conf/app.ini")

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
