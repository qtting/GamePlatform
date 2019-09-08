package Model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ApiToken struct {
	Id        int    `gorm:"not null"`
	Token     string `gorm:"not null"`
	UserName  string `gorm:"not null"`
}

func (apiToken ApiToken) TableName() string {
	return "api_token"
}

func (apiToken *ApiToken) BeforeCreate(scope *gorm.Scope) error {
	//var a = uuid.NewV4()
	//fmt.Println(a)
	//_ = scope.SetColumn("Token", interface{}(a))
	apiToken.Token = uuid.NewV4().String()

	return nil
}
//
//func (apiToken *ApiToken) AddTime(remember bool) {
//	now := time.Now()
//	duration, _ := time.ParseDuration("30m")
//	if remember {
//		apiToken.ExpiredAt = now.AddDate(0, 1, 0)
//	} else {
//		apiToken.ExpiredAt = now.Add(duration)
//	}
//}
