package api

var GroupAppApi = new(GroupApi)

type GroupApi struct {
	UserApi         // 匿名嵌套（嵌入）
	DemoApi DemoApi // 具名嵌套 那么 api.GroupApi{} 这种方式就拿不到值，需要api.GroupApi.DemoApi{}
}
