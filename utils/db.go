package utils

import (
	"github.com/jinzhu/gorm"
	"time"
)

var globalDB *gorm.DB

func InitDB() error {
	user := GetConf().String("db::user")
	password := GetConf().String("db::password")
	host := GetConf().String("db::host")
	port := GetConf().String("db::port")
	dbName := GetConf().String("db::dbname")
	parameters := GetConf().String("db::parameters")

	mysqlUrl := DSN(user, password, host, port, dbName, parameters)
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}

	//连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	globalDB = db

	return nil
}

func GetDB() *gorm.DB {
	return globalDB
}
