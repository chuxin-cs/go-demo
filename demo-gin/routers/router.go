package routers

import (
	"demo-gin/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 用户相关
	user := r.Group("/user")
	{
		user.GET("/list/:id", api.GetUser)
		user.GET("/list", api.GetUserList)
		//新建用户
		user.POST("/add", api.AddUser)
		//更新用户
		user.POST("/update", api.EditUser)
		//删除用户
		user.POST("/delete", api.DeleteUser)
	}
	err := r.Run(":9000")
	if err != nil {
		return nil
	}
	return r
}
