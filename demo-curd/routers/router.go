package routers

import (
	"github.com/chuxin-cs/demo-curd/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 用户相关
	user := r.Group("/user")
	{
		//获取用户列表
		user.GET("/list", api.GetUserList)
		//新建用户
		user.POST("/add", api.AddUser)
		//更新用户
		user.POST("/edit", api.EditUser)
		//删除用户
		user.POST("/delete", api.DeleteUser)
	}

	r.Run(":9000")

	return r
}
