package example

import (
	"github.com/gin-gonic/gin"
)

type DemoRouter struct {
}

func (d *DemoRouter) InitDemoRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	router := Router.Group("/demo")
	{
		router.GET("/list", DemoApi.GetList)
	}
}
