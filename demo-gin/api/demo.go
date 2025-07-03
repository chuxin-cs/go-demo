package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	// r.GET("/user/:id", func(){  }) 取前端传的参数
	id := c.Query("id")

	c.String(200, "get请求"+id)
}

func GetUserQueryList(c *gin.Context) {
	// 取前端传的参数
	//page1 := c.Query("page")
	//pageSize1 := c.Query("pageSize")
	// 取前端传的参数，没有的情况下用默认值
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	c.String(200, "get请求"+page+pageSize)
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(200, "post x-www-form-urlencoded请求"+name)
}

type Person struct {
	// json格式从name取值，并且该值为必须的
	Name string `json:"name" binding:"required"`
}

func EditUser(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, "post json格式请求")
}
