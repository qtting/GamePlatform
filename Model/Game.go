package Model

import "github.com/jinzhu/gorm"

//import "github.com/jinzhu/gorm"
//
//type Game struct {
//	game_name string `gorm:"primary_key;unique"`
//	introduction string `gorm:"unique;not null;default "`
//	download_times int `gorm:"not null"`
//}
type Game struct {
	GameName      string `gorm:"primary_key"`
	Version       string `gorm:"not null"`
	DownloadTimes int    `gorm:"not null"`
	Url           string `gorm:"not null"`
}

func (game *Game) TableName(scope *gorm.Scope) string {
	return "game"
}
