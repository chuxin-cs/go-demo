// 指明当前文件属于哪个包
package main

import "github.com/chuxin-cs/demo-curd/routers"

// 创建 main 方法
func main() {
	routers.InitRouter()
}
