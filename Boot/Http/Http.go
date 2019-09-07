package Http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Config"
)

var http *Http
var Router *gin.Engine

type Http struct {
	server *gin.Engine
	port string
	addr string
}

func InitHttp() {
	http = new(Http)
	http.server = gin.Default()
	Router = http.server
	http.port = Config.GetStringWithDefault("http.port", "8888")
	http.addr = Config.GetStringWithDefault("http.server", "localhost")
}

func Run() {
	err := http.server.Run(http.addr + ":" + http.port)
	if err != nil {
		panic(fmt.Errorf("Fatal error run http server: %s \n", err))
	}
}