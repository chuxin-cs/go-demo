package main

import (
	"github.com/gin-gonic/gin"
)

// 增
func add(c *gin.Context) {
	c.String(200, "add")
}

// 删
func del(c *gin.Context) {
	c.String(200, "del")
}

// 改
func edit(c *gin.Context) {
	c.String(200, "edit")
}

// 查
func getList(c *gin.Context) {
	c.String(200, "getList")
}

func main() {
	r := gin.Default()
	r.GET("/add", add)
	r.GET("/del", del)
	r.GET("/edit", edit)
	r.GET("/getList", getList)
	// 运行在 9000 端口
	r.Run(":9000")
}
