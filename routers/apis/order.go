package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "i am getList",
		"data": 1,
	})
}

func Add(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "i am add",
		"data": id,
	})
}
