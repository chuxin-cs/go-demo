package api

import "github.com/gin-gonic/gin"

type DemoApi struct {
}

func (m *DemoApi) GetList(c *gin.Context) {
	c.JSON(200, "我是user.go文件中的GitList")
}
