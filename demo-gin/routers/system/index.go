package system

import "demo-gin/api"

type GroupRouter struct {
	UserRouter
}

var (
	UserApi = api.GroupAppApi
)
