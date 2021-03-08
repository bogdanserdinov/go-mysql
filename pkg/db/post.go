package db

import (
	"awesomeProject/nix/entity"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"sync"
)

//WriteToDBPost - function required for writing data to post entity
func WriteToDBPost(p entity.Post,db *gorm.DB,mutex *sync.Mutex){
	mutex.Lock()
	//db.Create(&p)
	db.Select("UserID", "ID", "Title","Body").Create(&p)
	mutex.Unlock()
}
