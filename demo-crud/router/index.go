package router

import "github.com/gin-gonic/gin"

// Add 增
func Add(c *gin.Context) {
	c.String(200, "add")
}

// Del 删
func Del(c *gin.Context) {
	c.String(200, "del")
}

// Edit 改
func Edit(c *gin.Context) {
	c.String(200, "edit")
}

// GetList 查
func GetList(c *gin.Context) {
	c.String(200, "getList")
}
