package main

import (
	"github.com/chuxin-cs/go-demo/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":9999")
}
