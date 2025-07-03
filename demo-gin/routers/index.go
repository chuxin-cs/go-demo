package routers

import "demo-gin/api"

// 基于全局变量的特性
var (
	UserApi = api.GroupAppApi
	DemoApi = api.GroupAppApi.DemoApi
)
