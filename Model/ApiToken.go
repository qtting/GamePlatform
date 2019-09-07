package Model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ApiToken struct {
	gorm.Model

	Id        int    `gorm:"not null"`
	Token     string `gorm:"not null"`
	UserName  string `gorm:"not null"`
}

func (apiToken ApiToken) TableName() string {
	return "api_token"
}

func (apiToken *ApiToken) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("Token", interface{}(uuid.NewV4()))
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
