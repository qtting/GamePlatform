package Game

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Log"
	"github.com/qiankaihua/ginDemo/Boot/Orm"
	"github.com/qiankaihua/ginDemo/Model"
	"io"
	"os"
)


func FileUp(c *gin.Context) {
	form,_:= c.MultipartForm()
	var game Model.Game
	game.GameName = form.Value["name"][0]
	game.Version = form.Value["version"][0]

	fmt.Println(game.GameName)
	fmt.Println(form.Value["version"])

	file :=form.File["file"][0]

	filetwo,err := file.Open()
	defer filetwo.Close()
	if err != nil {
		Log.Info(err.Error())
	}

	fileName := game.GameName + game.Version


	out , err := os.Create(fileName)
	defer out.Close()

	if err != nil{
		Log.Info(err.Error())
	}

	_,err = io.Copy(out,filetwo)
	if err != nil{
		Log.Info(err.Error())
	}

	db := Orm.GetDB()
	if !db.Where("game_name = ? AND version = ?",game.GameName,game.Version).First(&game).RecordNotFound() {
		c.JSON(400,gin.H{"err_msg":"请换一个游戏名或者版本号"})
		return
	}

	fileUrl:="https://www.google.com/search?q="+fileName

	game.Url = fileUrl
	//if err := c.ShouldBindJSON(&data);err != nil {
	//	c.JSON(http.StatusBadRequest,gin.H{"err_msg":"格式不正确哦"})
	//	Log.Info(err.Error())
	//	return
	//}

	db.Create(&game)

}
