package Game

import (
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Orm"
	"github.com/qiankaihua/ginDemo/Model"
	"net/http"
)

//type download struct {
//	name string ``
//	version string ``
//}

func Download(c *gin.Context){

	name:=c.Param("name")
	version:=c.Param("version")

	var game Model.Game
	db := Orm.GetDB()
	db.Where("game_name = ? AND version = ?",name,version).First(&game)

	fileUrl:=game.Url

	db.Model(&game).Update("download_times",game.DownloadTimes+1)
	c.Redirect(http.StatusMovedPermanently,fileUrl)

}

