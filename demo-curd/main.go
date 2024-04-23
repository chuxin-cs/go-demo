// 指明当前文件属于哪个包
package main

// 导入 gin 包
import (
	"github.com/gin-gonic/gin"
)

// 创建 main 方法
func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// 打印 “你好，世界”
		c.String(200, "你好，世界")
	})

	r.Run(":9000")
}
