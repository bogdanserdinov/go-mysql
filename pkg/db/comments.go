package db

import (
	"awesomeProject/nix/entity"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

//WriteToDBComment - function required for writing data to comment entity
func WriteToDBComment(c entity.Comment,db *sql.DB,mutex *sync.Mutex) {
	mutex.Lock()

	_,err := db.Exec("insert into public.comments values(?,?,?,?,?) ",c.PostID,c.ID,c.Name,c.Email,c.Body)
	if err != nil {
		log.Fatal("failed in executing command to comments db : ",err)
	}

	mutex.Unlock()

}