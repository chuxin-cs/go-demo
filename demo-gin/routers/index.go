package routers

import (
	"demo-gin/routers/example"
	"demo-gin/routers/system"
)

var GroupRouterApp = new(GroupRouter)

type GroupRouter struct {
	System  system.GroupRouter
	Example example.GroupRouter
}
