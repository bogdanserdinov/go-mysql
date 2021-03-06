package db

import (
	"github.com/bogdanserdinov/go-mysql/entity"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"sync"
)

//WriteToDBComment - function required for writing data to comment entity
func WriteToDBComment(c entity.Comment, db *gorm.DB, mutex *sync.Mutex) {
	mutex.Lock()
	db.Create(&c)
	mutex.Unlock()
}
