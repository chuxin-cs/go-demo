// 指明当前文件属于哪个包
package main

// 导入 gin 包
import (
	"demo-gin/core"
)

// 创建 main 方法
func main() {
	core.InitRouter()
}
