package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser http://localhost:9000/user/list/123
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"msg": "传参是" + id,
	})
}

// GetUserList http://localhost:9000/user/list?page=1&pageSize=2
func GetUserList(c *gin.Context) {
	// 取前端传的参数
	page1 := c.Query("page")
	pageSize1 := c.Query("pageSize")
	// 取前端传的参数，没有的情况下用默认值
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	c.JSON(200, gin.H{
		"msg":  "如果不传那么默认的值：" + page + "-" + pageSize,
		"msg1": "取前端传递的参数：" + page1 + "-" + pageSize1,
	})
}

// AddUser http://localhost:9000/user/add
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	defaultValue := c.DefaultPostForm("value", "我是value的默认值，只有当我没传时才会是当前值")
	data := map[string]interface{}{
		"name":  name,
		"value": defaultValue,
	}
	c.JSON(200, gin.H{
		"data": data,
	})
}

type Person struct {
	// json格式从name取值，并且该值为必须的
	Name string `json:"name" binding:"required"`
}

// EditUser http://localhost:9000/user/update
func EditUser(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"person": person,
	})
}

// DeleteUser http://localhost:9000/user/delete?ids=1,2,3
func DeleteUser(c *gin.Context) {
	// 用 Query 或者 Param
	idsStr := c.Query("ids")
	c.JSON(200, gin.H{
		"msg": "删除用户",
		"ids": idsStr,
	})
}
