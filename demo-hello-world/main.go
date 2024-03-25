package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter() *gin.Engine {
	router := gin.New()
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/getList", getList)
	}
	return router
}
func getList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "i am getList",
		"data": 1,
	})
}
func main() {
	router := initRouter()
	router.Run(":9999")
}
