// 指明当前文件属于哪个包
package main

// 导入 gin 包
import "github.com/gin-gonic/gin"

// 创建 main 方法
func main() {
	// 在前端中 这一步操作好比 new Koa()
	r := gin.Default()
	// 定义了一个 GET 请求
	r.GET("/hello", func(c *gin.Context) {
		// 打印 “你好，世界”
		c.String(200, "你好，世界")
	})
	// 运行在9000端口
	r.Run(":9000")
}
