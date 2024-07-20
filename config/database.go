package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_cuti?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return DB
}
