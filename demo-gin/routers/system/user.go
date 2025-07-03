package system

import (
	"demo-gin/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (d *UserRouter) InitUserRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	UserApi1 := api.GroupAppApi.UserApi
	UserApi2 := api.UserApi{}
	UserApi3 := api.GroupApi{}
	router := Router.Group("/user")
	{
		router.GET("/list/:id", UserApi1.GetUser)
		router.GET("/list", UserApi2.GetUserList)
		//新建用户
		router.POST("/add", UserApi3.AddUser)
		//更新用户
		router.POST("/update", UserApi.EditUser)
		//删除用户
		router.POST("/delete", UserApi.DeleteUser)
	}
}
