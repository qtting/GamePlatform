package Game

import (
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Orm"
	"github.com/qiankaihua/ginDemo/Model"
)

func List(c *gin.Context){
	db := Orm.GetDB()
	var GameInfo []Model.Game
	db.Find(&GameInfo)
	c.JSON(200,gin.H{"data":GameInfo})
}
