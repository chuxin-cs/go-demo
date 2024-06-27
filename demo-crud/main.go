package main

import (
	"github.com/chuxin-cs/demo-crud/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/add", router.Add)
	r.GET("/del", router.Del)
	r.GET("/edit", router.Edit)
	r.GET("/getList", router.GetList)
	// 运行在 9000 端口
	r.Run(":9000")
}
