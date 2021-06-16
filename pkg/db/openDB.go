package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func OpenDataBase() *gorm.DB {
	dsn := "root:pass@tcp(127.0.0.1:3306)/public?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed in opening db : ", err)
	}
	return gormDB
}
