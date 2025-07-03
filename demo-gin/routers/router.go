package routers

import (
	"demo-gin/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 用户相关
	UserApi1 := api.GroupAppApi.UserApi
	UserApi2 := api.UserApi{}
	UserApi3 := api.GroupApi{}
	user := r.Group("/user")
	{
		user.GET("/list/:id", UserApi1.GetUser)
		user.GET("/list", UserApi2.GetUserList)
		//新建用户
		user.POST("/add", UserApi3.AddUser)
		//更新用户
		user.POST("/update", UserApi.EditUser)
		//删除用户
		user.POST("/delete", UserApi.DeleteUser)
	}

	r.GET("/demo", DemoApi.GetList)
	err := r.Run(":9000")
	if err != nil {
		return nil
	}
	return r
}
