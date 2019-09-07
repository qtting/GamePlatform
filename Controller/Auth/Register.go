package Auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Orm"
	"github.com/qiankaihua/ginDemo/Controller"
	"github.com/qiankaihua/ginDemo/Model"
	"net/http"
)

type RegisterValidate struct {
	UserName string `json:"username" binding:"required,len=20"`
	Password string `json:"password" binding:"required,len=40"`
}
func Register(c *gin.Context){
	var data RegisterValidate
	if err := c.ShouldBindJSON(&data);err != nil{
		fmt.Println(data.UserName)
		fmt.Println(data.Password)
		c.JSON(http.StatusBadRequest,gin.H{"message":"格式不正确哦"})
		return
	}
	var user Model.User
	db := Orm.GetDB()
	if db.Where("username = ?",data.UserName).First(&user).RecordNotFound(){
		user.Username = data.UserName
		user.Password = Controller.Sha256Get(data.Password)
		db.Create(&user)
		c.JSON(200,gin.H{"message":"注册成功"})
	} else{
		c.JSON(401,gin.H{"message":"请换一个用户名"})
	}
}




