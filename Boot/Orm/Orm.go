package Orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qiankaihua/ginDemo/Boot/Config"
)

var db *gorm.DB

func InitOrm() {
	engine := Config.GetString("database.engine")
	dbEngine, err := gorm.Open(engine, getParams(engine))
	if err != nil {
		panic(fmt.Errorf("Fatal error open database error [err=%s]!\n", err))
	}
	db = dbEngine
}

func GetDB() *gorm.DB {
	return db
}

func EndOrm() {
	err := db.Close()
	if err != nil {
		panic(fmt.Errorf("Fatal error close database error [err=%s]!\n", err))
	}
}

func getParams(engine string) string {
	switch engine {
	case "mysql":
		host := Config.GetStringWithDefault("database.host", "localhost")
		port := Config.GetStringWithDefault("database.port", "3306")
		dbname := Config.GetStringWithDefault("database.dbname", "GamePlatform")
		username := Config.GetStringWithDefault("database.user", "root")
		password := Config.GetStringWithDefault("database.password", "")
		mysqlParams := Config.GetStringWithDefault(
			"database.mysqlParams",
			"parseTime=True&charset=utf8mb4&loc=Local")
		params := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
			username, password, host, port, dbname, mysqlParams)
		return params
	case "sqlite3":
		params := Config.GetStringWithDefault("database.dbname", "/tmp/gorm.db")
		return params
	case "postgres":
		host := Config.GetStringWithDefault("database.host", "localhost")
		port := Config.GetStringWithDefault("database.port", "5432")
		dbname := Config.GetStringWithDefault("database.dbname", "QKteam")
		username := Config.GetStringWithDefault("database.user", "root")
		password := Config.GetStringWithDefault("database.password", "")
		sslMode := Config.GetStringWithDefault("database.sslmode", "disable")
		params := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, username, password, dbname, sslMode)
		return params
	default:
		panic(fmt.Errorf("Fatal error getting database params: [engine=%s] \n", engine))
	}
}
