package Route

import (
	"github.com/qiankaihua/ginDemo/Boot/Http"
)

// 静态文件路由转发请写在这个文件中
func AddStaticRoute() {
	Http.Router.Static("/static", "./Public")
}