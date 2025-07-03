package example

import "demo-gin/api"

type GroupRouter struct {
	DemoRouter
}

// DemoApi 基于全局变量的特性
var (
	DemoApi = api.GroupAppApi.DemoApi
)
