package Model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (user *User) TableName(scope *gorm.Scope) string {
	return "user"
}

//func (user *User) GetData(kind string) map[string]interface{} {
//	switch kind {
//	case "detail":
//		return map[string]interface{}{
//			"id":       user.ID,
//			"username": user.Username,
//		}
//	case "profile":
//		return map[string]interface{}{
//			"username": user.Username,
//		}
//
//	default:
//		return make(map[string]interface{})
//	}
//}
