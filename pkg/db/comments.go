package db

import (
	"awesomeProject/nix/entity"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"sync"
)

//WriteToDBComment - function required for writing data to comment entity
func WriteToDBComment(c entity.Comment,db *gorm.DB,mutex *sync.Mutex) {
	mutex.Lock()
	//db.Create(&c)
	db.Create(&c)
	mutex.Unlock()

}