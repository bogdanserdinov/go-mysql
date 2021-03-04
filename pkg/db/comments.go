package db

import (
	"awesomeProject/nix/entity"
	"database/sql"
	"log"
	"sync"
)

func WriteToDBComment(c entity.Comment,mutex *sync.Mutex) {
	mutex.Lock()
	db,err := sql.Open("mysql","root:blablabla29032002@/practice")

	if err != nil {
		log.Fatal("failed in opening db : ",err)
	}

	_,err = db.Exec("insert into practice.comments values(?,?,?,?,?) ",c.PostID,c.ID,c.Name,c.Email,c.Body)

	if err != nil {
		log.Fatal("failed in executing command to comments db : ",err)
	}

	if err = db.Close();err != nil{
		log.Fatal("failed to close db : ",err)
	}
	mutex.Unlock()

}