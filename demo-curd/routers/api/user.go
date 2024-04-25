package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	name := c.DefaultQuery("name", "chuxin")
	c.String(http.StatusOK, "获取用户列表"+name)
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"name":   name,
	})
}

type Person struct {
	Name string `json:"name" binding:"required"` // json格式从name取值，并且该值为必须的
}

func EditUser(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var name = person.Name

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"name":    name,
	})
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "删除用户")
}
