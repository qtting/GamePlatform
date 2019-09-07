package Route

import (
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Http"
	"github.com/qiankaihua/ginDemo/Controller/Auth"
	//"github.com/qiankaihua/ginDemo/Controller/Game"
)

// 该文件编写 api 路由转发
func AddApiRoute() {
	Http.Router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "success",
			"version": "v1.0",
		})
	})

	auth := Http.Router.Group("/auth")
	{
		auth.POST("register", Auth.Register)

		auth.POST("login", Auth.Login)
	}

	//gameload := Http.Router.Group("/game")
	{
		//gameload.GET("download",Game.Download)
		//
		//gameload.POST("upload",Game.Upload)
	}
}