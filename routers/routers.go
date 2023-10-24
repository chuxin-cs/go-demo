package routers

import (
	v1 "github.com/chuxin-cs/go-demo/routers/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// gin 日志
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/getList", v1.GetList)
		apiV1.GET("/add/:id", v1.Add)
	}

	// 返回实例
	return router
}
