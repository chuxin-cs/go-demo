package core

import (
	"demo-gin/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 实例化
	Router := gin.New()
	// 创建 公开
	PublicGroup := Router.Group("/api")
	// 创建 私有
	PrivateGroup := Router.Group("/api")

	// 开始注入
	{
		routers.GroupRouterApp.System.UserRouter.InitUserRouter(PrivateGroup, PublicGroup)
		routers.GroupRouterApp.Example.DemoRouter.InitDemoRouter(PrivateGroup, PublicGroup)
	}

	err := Router.Run(":9000")
	if err != nil {
		return nil
	}
	return Router
}
