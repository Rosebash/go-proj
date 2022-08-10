package config

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "rosebash:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
