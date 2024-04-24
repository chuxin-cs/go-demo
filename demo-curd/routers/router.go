package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// 打印 “你好，世界”
		c.String(200, "你好，世界")
	})

	r.Run(":9000")

	return r
}
